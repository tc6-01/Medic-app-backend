package Handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegitserHandler 处理注册请求
func RegitserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接受请求对象
		var requestBody struct {
			Username      string `json:"username"`
			Password      string `json:"password"`
			RetryPassword string `json:"retry_password"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			handleError(c, "param  not found", "")
			return
		}
		fmt.Println(requestBody)
		if requestBody.Password != requestBody.RetryPassword {
			handleError(c, "register error", "password not match")
			return
		}
		if len(requestBody.Username) == 0 || len(requestBody.Password) == 0 {
			handleError(c, "username or password not found", "username or password not found")
			return
		}
		if len(requestBody.Username) < 6 {
			handleError(c, "register error", "username length must be greater than 6")
			return
		}
		if len(requestBody.Username) > 16 {
			handleError(c, "register error", "username length must be less than 16")
			return
		}
		if len(requestBody.Password) < 6 {
			handleError(c, "register error", "password length must be greater than 8")
			return
		}
		err := RegisterUser(db, requestBody.Username, requestBody.Password)
		if err != nil {
			handleError(c, err.Error(), "Register Failed")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "Register Successfully"})
	}
}
