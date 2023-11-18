package Handler

import (
	"database/sql"
	"net/http"
	"yiliao/Dao"

	"github.com/gin-gonic/gin"
)

// 返回当前用户所能查看的PDF
func GetFileListHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取用户信息（例如，从解析的 JWT 中获取用户 ID）
		// 如果你的用户身份验证逻辑已经在其他地方完成，可以从上下文中获取用户信息
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "has token but no user found!")
			return
		}

		// 查询数据库以获取指定用户的所有 PDF 文件
		query := "SELECT expire, file_name, file_size, owner, use_count, use_limit FROM file_list WHERE owner = ? UNION SELECT expire, file_name, file_size, owner, use_count, use_limit FROM file_share_list WHERE target = ?"
		rows, err := db.Query(query, user, user)
		if err != nil {
			handleError(c, "DB Error", "querry error !")
			return
		}
		defer rows.Close()

		// 构建文件列表
		var fileList []Dao.FileListElement // 请确保 FileListElement 结构体与数据库表的列对应
		for rows.Next() {
			var file Dao.FileListElement
			err := rows.Scan(&file.Expire, &file.FileName, &file.FileSize, &file.Owner, &file.Use_count, &file.UseLimit)
			if err != nil {
				handleError(c, "Parser Error", "parser data to sturct error!")
				return
			}
			fileList = append(fileList, file)
		}

		// 检查是否有错误发生
		if err := rows.Err(); err != nil {
			handleError(c, "DB Error", "rows error!")
			return
		}

		// 返回文件列表
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Success", "data": fileList})
	}
}
