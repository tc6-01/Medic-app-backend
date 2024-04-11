package Handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"yiliao/Dao"
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
		// 验证用户登录
		token, err := authenticateUser(db, requestBody.Username, requestBody.Password)
		if err != nil {
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05")
			// 登陆失败更新失败用户表
			_, err := db.Exec(`insert into t_login_fail(user_name,e_time) values(?,?)`, requestBody.Username, formattedTime)
			if err != nil {
				handleError(c, err.Error(), "Login failed!")
				return
			}
		}
		var rowId int
		err = db.QueryRow(`select role_id from user where user_name = ?`, requestBody.Username).Scan(&rowId)
		if err != nil {
			return
		}
		// 记录当前登录态
		c.Set("isAdmin", rowId)
		log.Println(c.Get("isAdmin"))
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK",
			"data": gin.H{"token": token, "isAdmin": rowId}})
	}
}

// RegisterHandler 处理注册请求
func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接受请求对象
		var requestBody struct {
			Username      string `json:"username"`
			Password      string `json:"password"`
			Role          string `json:"role"`
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
		err := RegisterUser(db, requestBody.Username, requestBody.Password, requestBody.Role)
		if err != nil {
			handleError(c, err.Error(), "Register Failed")
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "注册成功"})
	}
}

// GetUsersHandler 返回所有用户名
func GetUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取当前用户
		user, _ := c.Get("user")
		// 查询数据库中的所有用户
		rows, err := db.Query("SELECT user_name FROM user")
		if err != nil {
			handleError(c, "DB Error", "查询失败")
			return
		}
		defer rows.Close()
		d := user.(*Dao.User)
		var usernames []string
		// 遍历查询结果，将用户名添加到切片中
		for rows.Next() {
			var username string
			err := rows.Scan(&username)
			if err != nil {
				handleError(c, "Parser Error", "请求失败")
				return
			}
			if d.Username == username {
				continue
			}
			usernames = append(usernames, username)
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功获取用户列表", "data": usernames})
	}
}
