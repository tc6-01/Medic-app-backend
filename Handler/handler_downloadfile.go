package Handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "user not found")
			return
		}

		// 查询数据库以检查用户是否拥有指定文件   1. file_list用户拥有该文件  2. file_share_list用户被分享了该文件
		query := "SELECT owner, file_name FROM file_list WHERE owner = ? AND file_name = ? AND use_count< use_limit union SELECT owner, file_name FROM file_share_list WHERE target = ? AND file_name = ? AND use_count< use_limit"
		var owner string //记录一下这个文件的拥有者，根据 owner 和user 是否相同来选择更新两个数据表之一。
		var fileName string
		err := db.QueryRow(query, user, filename, user, filename).Scan(&owner, &fileName)
		if err != nil {
			if err == sql.ErrNoRows {
				handleError(c, "DB Error", "File not found or limit found")
				return
			}
			handleError(c, "DB Error", "mysql error")
			return
		}
		//更新次数
		if user == owner { //如果下载人是文件所有者，则更新 file_list表
			_, err = db.Exec(`UPDATE file_list SET use_count = use_count + 1 WHERE owner = ? AND file_name = ?`,
				user, fileName)
			if err != nil {
				handleError(c, "DB Error", "error update count")
			}
		} else { //不是文件拥有者，则更新 fiel_share_list
			_, err = db.Exec(`UPDATE file_share_list SET use_count = use_count + 1 WHERE target = ? AND file_name = ?`,
				user, fileName)
			if err != nil {
				handleError(c, "DB Error", "error update count")
			}
		}
		//

		// 从当前目录下读取文件
		filePath := fileName // 假设文件名与当前目录下的文件路径对应
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading file"})
			return
		}
		// 设置响应头，指定文件名和内容类型
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		// 将文件数据发送给客户端
		c.Data(http.StatusOK, "application/pdf", fileData)
	}
}
