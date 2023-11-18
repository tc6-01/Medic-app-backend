package Handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsersHandler 返回所有用户名
func GetUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//权限验证
		// 从请求头中获取Authorization标头
		authHeader := c.GetHeader("Authorization")

		// 检查Authorization标头是否存在
		if authHeader == "" {
			handleError(c, "Unauthorized", "Authorization dosen't exit")
			return
		}
		//获取自己的名字
		user, exists := c.Get("user")
		if !exists {
			handleError(c, "Unauthorized", "has token but no user found!")
			return
		}
		// 查询数据库中的所有用户
		rows, err := db.Query("SELECT username FROM users WHERE username!= ?", user)
		if err != nil {
			handleError(c, "DB Error", "querry error !")
			return
		}
		defer rows.Close()

		var usernames []string

		// 遍历查询结果，将用户名添加到切片中
		for rows.Next() {
			var username string
			err := rows.Scan(&username)
			if err != nil {
				handleError(c, "Parser Error", "parser data to sturct error!")
				return
			}
			usernames = append(usernames, username)
		}
		//
		// 返回用户名列表
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Success", "data": usernames})
	}
}
