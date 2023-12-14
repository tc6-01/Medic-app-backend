package Handler

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"yiliao/Dao"
	"yiliao/Utils"

	"github.com/gin-gonic/gin"
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
	err := db.QueryRow("SELECT user_id, password, role_id, public_key FROM user WHERE user_name = ?", username).Scan(
		&userId, &storedPassword, &roleId, &ak)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	// 使用盐值对密码进行加密
	password += "yiliao"
	hashedPassword := cryptoPass(password)
	if err != nil {
		log.Println("密码加密失败")
		return "", err
	}
	if hashedPassword != storedPassword {
		log.Println("密码错误")
		return "", fmt.Errorf("incorrect password")
	}
	user := Dao.User{
		UserId:    userId,
		Username:  username,
		Role:      roleId,
		PublicKey: ak,
	}
	log.Println(user.UserId)
	return Utils.CreateToken(user), nil
}

// 用户注册
func RegisterUser(db *sql.DB, username, password string) error {
	userId := 0
	err := db.QueryRow("SELECT user_id FROM user WHERE user_name = ?", username).Scan(
		&userId)
	// 如果成功查询说明用户已存在
	if err == nil {
		return fmt.Errorf("user already exeists!")
	}
	password += "yiliao"
	hashedPassword := cryptoPass(password)
	// 向数据库中插入用户
	_, err = db.Exec("insert into user(user_name,password) values(?,?)", username, string(hashedPassword))
	if err != nil {
		log.Println("插入数据库失败")
		return err
	}
	return nil
}

func cryptoPass(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 根据字符串错误来处理
func handleError(c *gin.Context, errors, msg string) {
	switch errors {
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
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "msg": errors, "data": "[]"})
	}
}
