package Handler

import (
	"net/http"
	"strings"
	"yiliao/Utils"

	"github.com/gin-gonic/gin"
)

func RouterAuth(c *gin.Context) {
	// 获取路由路径
	path := c.FullPath()
	// 如果路由是登录路径，则不执行验证
	if path == "/user/login" || path == "/user/register" {
		c.Next()
		return
	}

	authHeader := c.GetHeader("Authorization")
	//没有authHeader
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "用户未登录", "data": ""})
		c.Abort() // 停止请求处理
		return
	}

	parts := strings.Split(authHeader, " ")
	//authHeader格式不对
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "用户未登录", "data": ""})
		c.Abort() // 停止请求处理
		return
	}

	tokenString := parts[1]
	user, err := Utils.ValidateToken(tokenString)
	//token 不对
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 500, "msg": "用户未登录", "data": ""})
		c.Abort() // 停止请求处理
		return
	}
	c.Set("user", user)
	// 记录当前登录态
	c.Set("isAdmin", user.Role)
	c.Next() // 继续请求处理
}
