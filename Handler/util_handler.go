package Handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"yiliao/Dao"
	"yiliao/Utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 验证用户登录
func authenticateUser(db *sql.DB, username, password string) (string, error) {
	var (
		storedPassword string
		userId         int64
		roleId         int
		ak             string
	)
	/*
		从数据库中查询获取用户信息，并构建用户对象
	*/
	err := db.QueryRow("SELECT user_id, password, role_id, public_key FROM users WHERE user_name = ?", username).Scan(
		&userId, &storedPassword, &roleId, &ak)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	// 使用盐值对密码进行加密
	password += "yiliao"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	if string(hashedPassword) != storedPassword {
		return "", fmt.Errorf("incorrect password")
	}
	user := Dao.User{
		UserId:    userId,
		Username:  username,
		Role:      roleId,
		PublicKey: ak,
	}
	return Utils.CreateToken(user), nil
}

// 用户注册
func RegisterUser(db *sql.DB, username, password string) error {
	var userId int64
	err := db.QueryRow("SELECT user_id, FROM users WHERE user_name = ?", username).Scan(
		&userId)
	if err != nil {
		return fmt.Errorf("user already exeists!")
	}
	password += "yiliao"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 向数据库中插入用户
	_, err = db.Exec("insert into user(user_name,password) values(?,?)", username, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// 根据字符串错误来处理
func handleError(c *gin.Context, handleError, msg string) {
	switch handleError {
	case "Login Error":
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": msg, "data": "[]"})
	case "Unauthorized":
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": msg, "data": "[]"})
	case "Register Error":
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": msg, "data": "[]"})
	case "DB Error":
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": msg, "data": "[]"})
	case "Parser Error":
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": msg, "data": "[]"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "msg": handleError, "data": "[]"})
	}
}
