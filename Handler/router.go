package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"yiliao/Utils"
)

func RouterAuth(c *gin.Context) {
	// 获取路由路径
	path := c.FullPath()

	// 如果路由是登录路径，则不执行验证
	if path == "/login" {
		c.Next()
		return
	}

	authHeader := c.GetHeader("Authorization")
	//没有authHeader
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "Missing Authorization header", "data": ""})
		c.Abort() // 停止请求处理
		return
	}

	parts := strings.Split(authHeader, " ")
	//authHeader格式不对
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "Invalid Authorization header format", "data": ""})
		c.Abort() // 停止请求处理
		return
	}

	tokenString := parts[1]
	user, err := Utils.ValidateToken(tokenString)
	//token 不对
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "Invalid token", "data": ""})
		c.Abort() // 停止请求处理
		return
	}

	c.Set("user", user.Username) // 存储用户信息到上下文

	c.Next() // 继续请求处理
}
