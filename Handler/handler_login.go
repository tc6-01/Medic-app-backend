package Handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler 处理登录请求
func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			handleError(c, "Login Error", "param  not found")
			return
		}
		log.Println(requestBody.Password)
		// 验证用户登录
		token, err := authenticateUser(db, requestBody.Username, requestBody.Password)
		if err != nil {
			handleError(c, err.Error(), "Login failed!")
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK", "data": gin.H{"token": token}})
	}
}
