package Handler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
	"yiliao/Dao"
)

// DownloadFileHandler 提供PDF文件预览功能
func DownloadFileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取参数
		filename := c.Query("filename")
		fmt.Println("got filename:", filename)
		if filename == "" {
			handleError(c, "Parser Error", "文件名传输错误")
			return
		}
		state := c.Query("state")
		fmt.Println("got state:", state)
		if !(state == "owned" || state == "shared" || state == "fromShared") {
			handleError(c, "Parser Error", "病历状态传输错误")
			return
		}
		// 拥有不需要更新次数
		// 从上下文中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		user, exist := c.Get("user")
		if !exist {
			handleError(c, "Unauthorized", "用户未登录")
			return
		}
		d := user.(*Dao.User)
		var id, useCount, useLimit int
		// 被共享需要直接减去当前记录的次数  当前共享需要减去上一次共享的用户的可访问次数
		if state == "shared" || state == "fromShared" {
			flag := true
			// 利用当前用户名和文件名找到共享记录,增加共享发起者ID识别，避免多文件返回失败
			query := `select id, use_count, use_limit from share_files where  target_user_id = ? and fileId =  
					(select file_id from files where file_name = ? )`
			err := db.QueryRow(query, d.UserId, filename).Scan(&id, &useCount, &useLimit)
			if err != nil {
				if errors.Is(sql.ErrNoRows, err) {
					log.Println("No need to update")
					flag = false
				} else {
					handleError(c, "DB Error", "数据库查询错误")
					return
				}
			}
			if flag {
				if useCount >= useLimit {
					handleError(c, "DB Error", "可访问次数已用完，请联系共享者")
					return
				}
				//更新次数
				_, err = db.Exec(`UPDATE share_files SET use_count = ? WHERE id = ?`, useCount+1, id)
				if err != nil {
					handleError(c, "DB Error", "更新访问次数错误")
					return
				}
			}

		}
		// 从当前目录下读取文件
		fileData, err := os.ReadFile("WebServer/files/" + filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "文件不存在"})
			return
		}
		// 设置响应头，指定文件名和内容类型
		c.Header("Content-Disposition", "attachment; filename="+filename)
		// 将文件数据发送给客户端
		c.Data(http.StatusOK, "application/pdf", fileData)
	}
}

// UploadFile 上传文件
func UploadFile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证当前用户为管理员
		root, exist := c.Get("isAdmin")
		if !exist || root == 0 {
			handleError(c, "Unauthorized", "非管理员不可进行文件上传")
			return
		}
		uName := c.PostForm("owner")
		file, err := c.FormFile("files")
		var sId int64
		// 识别文件后缀名并在之前加上用户名
		name := "(" + uName + ")" + file.Filename
		// 上传之后，需要添加表索引（增加文件共享记录）数据库进行更新
		err = db.QueryRow("SELECT user_id FROM user WHERE user_name = ?", uName).Scan(&sId)
		if err != nil {
			handleError(c, "DB Error", "该用户不存在")
			return
		}
		// 管理员上传文件使用默认共享策略（过期时间为2025年，最大使用100次限制）
		_, err = db.Exec("INSERT INTO files(file_name, file_size, owner_id) values (?,?,?)",
			name, file.Size, sId)
		if err != nil {
			handleError(c, "DB Error", "文件已存在")
			return
		}
		// 将病历下载到服务器
		if err := c.SaveUploadedFile(file, "WebServer/files/"+name); err != nil {
			fmt.Printf("SaveUploadedFile,err=%v", err)
			handleError(c, "Upload Error", "上传失败")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "文件已上传"})
	}
}

// GetFileListHandler 返回当前用户可查看的病历
func GetFileListHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		// 如果你的用户身份验证逻辑已经在其他地方完成，可以从上下文中获取用户信息
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "当前用户未登录")
			return
		}

		// 查询数据库以获取指定用户的所有 PDF 文件 并获取该文件对应的病例共享策略
		// 用户拥有的病历、被共享的病历
		d := user.(*Dao.User)
		log.Println("当前时间", time.Now().Format("2006-01-02 15:04:05"))
		rows, err := db.Query(`select owner_id, files.expire, file_name, file_size,use_limit,use_count,is_allow
			from files 
			where owner_id = ?
			union select owner_id, share_files.expire, file_name, file_size, share_files.use_limit, share_files.use_count,share_files.is_allow 
			from files 
			LEFT JOIN share_files ON files.file_id = share_files.fileId 
			where share_files.target_user_id = ? `, d.UserId, d.UserId)
		defer rows.Close()
		if err != nil {
			handleError(c, "DB Error", "查询失败")
			return
		}
		log.Println("结束时间", time.Now().Format("2006-01-02 15:04:05"))
		// 构建文件列表
		var fileList []Dao.FileListElement // 请确保 FileListElement 结构体与数据库表的列对应
		for rows.Next() {
			var file Dao.FileListElement
			var uid int64
			err := rows.Scan(&uid, &file.Expire, &file.FileName, &file.FileSize, &file.UseLimit, &file.UseCount, &file.IsAllow)
			if err != nil {
				handleError(c, "Parser Error", "查询失败")
				return
			}
			row, err := db.Query(`select user_name from user where user_id= ?`, uid)
			if err != nil {
				handleError(c, "DB Error", "查询失败")
				return
			}
			if row.Next() {
				err := row.Scan(&file.Owner)
				if err != nil {
					handleError(c, "Parser Error", "查询失败")
				}
			}
			row.Close()
			fileList = append(fileList, file)
		}
		// 返回文件列表
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Success", "data": fileList})
	}
}

// GetShareHandler 获取共享病历
func GetShareHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "当前用户未登录")
			return
		}
		d := user.(*Dao.User)
		// 执行 SQL 查询以获取当前用户共享的所有文件
		query := `
            SELECT share_files.id,files.file_name, files.file_size, user.user_name, share_files.expire, share_files.use_limit,share_files.use_count,share_files.is_allow
            FROM share_files 
        	left join files on share_files.fileId = files.file_id 
            left join user on share_files.target_user_id = user.user_id
            WHERE from_user_id = ?`
		rows, err := db.Query(query, d.UserId)
		defer rows.Close()
		if err != nil {
			handleError(c, "DB Error", "Database query error")
			return
		}
		// 创建一个切片来存储共享文件的信息
		var sharedFiles []Dao.ShareFile

		// 遍历查询结果并填充 sharedFiles 切片
		for rows.Next() {
			var (
				id       int64
				expire   int64
				fileName string
				target   string
				isAllow  int64
				useCount int64
				useLimit int64
				fileSize int64
			)

			err := rows.Scan(&id, &fileName, &fileSize, &target, &expire, &useLimit, &useCount, &isAllow)
			if err != nil {
				handleError(c, "DB Error", "Error scanning rows")
				return
			}

			// 创建共享文件对象并填充数据
			sharedFile := Dao.ShareFile{
				Id:       id,
				Expire:   expire,
				FileName: fileName,
				Target:   target,
				IsAllow:  isAllow,
				UseCount: useCount,
				UseLimit: useLimit,
				FileSize: fileSize,
			}

			sharedFiles = append(sharedFiles, sharedFile)
		}
		// 返回共享文件列表
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Success", "data": sharedFiles})
	}
}

// GetBeShareHandler  获取被共享病历
func GetBeShareHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "当前用户未登录")
			return
		}
		d := user.(*Dao.User)
		// 执行 SQL 查询以获取当前用户共享的所有文件
		query := `
			SELECT share_files.id, files.file_name, files.file_size, user.user_name, share_files.expire, share_files.use_limit,share_files.use_count,share_files.is_allow
			FROM share_files
			left join files on share_files.fileId = files.file_id
			left join user on share_files.from_user_id = user.user_id
			WHERE target_user_id = ?`
		rows, err := db.Query(query, d.UserId)
		defer rows.Close()
		if err != nil {
			handleError(c, "DB Error", "Database query error")
			return
		}
		// 创建一个切片来存储共享文件的信息
		var sharedFiles []Dao.BeShareFile

		// 遍历查询结果并填充 sharedFiles 切片
		for rows.Next() {
			var (
				id       int64
				expire   int64
				fileName string
				from     string
				useCount int64
				isAllow  int64
				useLimit int64
				size     int64
			)

			err := rows.Scan(&id, &fileName, &size, &from, &expire, &useLimit, &useCount, &isAllow)
			if err != nil {
				handleError(c, "DB Error", "Error scanning rows")
				return
			}

			// 创建共享文件对象并填充数据
			beSharedFile := Dao.BeShareFile{
				Id:       id,
				Expire:   expire,
				FileName: fileName,
				From:     from,
				IsAllow:  isAllow,
				UseCount: useCount,
				UseLimit: useLimit,
				FileSize: size,
			}

			sharedFiles = append(sharedFiles, beSharedFile)
		}
		// 返回共享文件列表
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Success", "data": sharedFiles})
	}
}

func DeleteShareFileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sId := c.Query("id")
		_, err := db.Exec("DELETE from share_files where id = ?", sId)
		if err != nil {
			handleError(c, "DB Error", "共享记录不存在")
			return
		}
		// 返回共享文件列表
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "共享记录删除成功", "data": ""})
	}
}
