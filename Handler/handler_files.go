package Handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
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
			handleError(c, "Parser Error", "filename not transfer")
			return
		}

		// 从上下文中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		// 如果你的用户身份验证逻辑已经在其他地方完成，可以从上下文中获取用户信息
		_, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "user not found")
			return
		}

		//// 查询数据库以检查用户是否拥有指定文件   1. file_list用户拥有该文件  2. file_share_list用户被分享了该文件
		//query := "SELECT owner, file_name FROM file_list WHERE owner = ? AND file_name = ? AND use_count< use_limit union SELECT owner, file_name FROM file_share_list WHERE target = ? AND file_name = ? AND use_count< use_limit"
		//var owner string //记录一下这个文件的拥有者，根据 owner 和user 是否相同来选择更新两个数据表之一。
		//var fileName string
		//err := db.QueryRow(query, user, filename, user, filename).Scan(&owner, &fileName)
		//if err != nil {
		//	if errors.Is(sql.ErrNoRows, err) {
		//		handleError(c, "DB Error", "File not found or limit found")
		//		return
		//	}
		//	handleError(c, "DB Error", "mysql error")
		//	return
		//}
		////更新次数
		//if user == owner { //如果下载人是文件所有者，则更新 file_list表
		//	_, err = db.Exec(`UPDATE file_list SET use_count = use_count + 1 WHERE owner = ? AND file_name = ?`,
		//		user, fileName)
		//	if err != nil {
		//		handleError(c, "DB Error", "error update count")
		//	}
		//} else { //不是文件拥有者，则更新 fiel_share_list
		//	_, err = db.Exec(`UPDATE file_share_list SET use_count = use_count + 1 WHERE target = ? AND file_name = ?`,
		//		user, fileName)
		//	if err != nil {
		//		handleError(c, "DB Error", "error update count")
		//	}
		//}
		//

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
		name := c.PostForm("fileName")
		sizeStr := c.PostForm("fileSize")
		oIdStr := c.PostForm("owner")
		file, err := c.FormFile("files")
		size, _ := strconv.Atoi(sizeStr)
		var sId int64
		// 上传之后，需要添加表索引（增加文件共享记录）数据库进行更新
		err = db.QueryRow("SELECT user_id FROM user WHERE user_name = ?", oIdStr).Scan(&sId)
		if err != nil {
			handleError(c, "DB Error", "该用户不存在")
			return
		}
		// 管理员上传文件使用默认共享策略（过期时间为2025年，最大使用100次限制）
		_, err = db.Exec("INSERT INTO files(file_name, file_size, owner_id) values (?,?,?)",
			name, size, sId)
		if err != nil {
			handleError(c, "DB Error", "文件已存在")
			return
		}
		// 将病历下载到服务器
		if err := c.SaveUploadedFile(file, "webServer/files/"+name); err != nil {
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
		// 用户拥有的病历、被共享的病历、共享的病例
		d := user.(*Dao.User)
		log.Println("当前时间", time.Now().Format("2006-01-02 15:04:05"))
		rows, err := db.Query(`select owner_id, files.expire, file_name, file_size,use_limit,use_count 
			from files 
			where owner_id = ?
			union select owner_id, share_files.expire, file_name, file_size, share_files.use_limit, share_files.use_count 
			from files 
			LEFT JOIN share_files ON files.file_id = share_files.fileId 
			where share_files.target_user_id = ? or share_files.from_user_id = ?`, d.UserId, d.UserId, d.UserId)
		if err != nil {
			handleError(c, "DB Error", "查询失败")
			return
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				handleError(c, "DB Error", "close error !")
				return
			}
		}(rows)
		log.Println("结束时间", time.Now().Format("2006-01-02 15:04:05"))
		// 构建文件列表
		var fileList []Dao.FileListElement // 请确保 FileListElement 结构体与数据库表的列对应
		for rows.Next() {
			var file Dao.FileListElement
			var uid int64
			err := rows.Scan(&uid, &file.Expire, &file.FileName, &file.FileSize, &file.UseLimit, &file.Use_count)
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
			fileList = append(fileList, file)
		}

		// 检查是否有错误发生
		if err := rows.Err(); err != nil {
			handleError(c, "DB Error", "rows error!")
			return
		}

		// 返回文件列表
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Success", "data": fileList})
	}
}

// GetShareHandler 获取当前用户共享的所有文件
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
		rows, err := db.Query(`
            SELECT files.file_name, files.file_size, user.user_name, share_files.expire, share_files.use_limit,share_files.use_count
            FROM share_files 
        	left join files on share_files.fileId = files.file_id 
            left join user on share_files.target_user_id = user.user_id
            WHERE from_user_id = ?`, d.UserId)

		if err != nil {
			handleError(c, "DB Error", "Database query error")
			return
		}
		defer rows.Close()

		// 创建一个切片来存储共享文件的信息
		var sharedFiles []Dao.ShareFile

		// 遍历查询结果并填充 sharedFiles 切片
		for rows.Next() {
			var (
				expire   int64
				fileName string
				target   string
				useCount int64
				useLimit int64
				fileSize int64
			)

			err := rows.Scan(&fileName, &fileSize, &target, &expire, &useLimit, &useCount)
			if err != nil {
				handleError(c, "DB Error", "Error scanning rows")
				return
			}

			// 创建共享文件对象并填充数据
			sharedFile := Dao.ShareFile{
				Expire:    expire,
				FileName:  fileName,
				Target:    target,
				Use_count: useCount,
				UseLimit:  useLimit,
				FileSize:  fileSize,
			}

			sharedFiles = append(sharedFiles, sharedFile)
		}

		// 检查是否有错误
		if err := rows.Err(); err != nil {
			handleError(c, "DB Error", "Error in rows")
			return
		}

		// 返回共享文件列表
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Success", "data": sharedFiles})
	}
}
