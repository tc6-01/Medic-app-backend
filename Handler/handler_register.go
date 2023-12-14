package Handler

import (
	"database/sql"
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
			handleError(c, "param  not found", "传递参数错误")
			return
		}
		if requestBody.Password != requestBody.RetryPassword {
			handleError(c, "Register Error", "password not match")
			return
		}
		err := RegisterUser(db, requestBody.Username, requestBody.Password)
		if err != nil {
			handleError(c, err.Error(), "Register Failed")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "注册成功"})
	}
}
