package Utils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"yiliao/Dao"
)

/*
jwt 创建Token、验证Token 登录
*/
var jwtKey = []byte("4295d277e69c8be76df37abbae2d989e9d8482dd")

// ToJson /*
func ToJson(user Dao.User) string {
	jsonData, _ := json.Marshal(user)
	return string(jsonData)
}

// ToUser /*
func ToUser(str string) *Dao.User {
	var user Dao.User
	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		return nil
	}
	return &user
}

// CreateToken 签发令牌
func CreateToken(user Dao.User) string {
	// 设置令牌的过期时间
	expirationTime := time.Now().Add(24 * time.Hour)

	/*
		JWT令牌中存放用户对象
		将用户对象封装成Json字符串
		使用Base64Url编码进行加密
	*/
	log.Println("获取用户信息中")
	userObj := ToJson(user)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "your-issuer", // 可以自定义
		Subject:   userObj,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("正在进行签名")
	// 使用密钥签名JWT令牌
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println("签名成功")
	return tokenString
}

// ValidateToken 验证JWT令牌  ,参数为String 返回值为User
func ValidateToken(tokenString string) (*Dao.User, error) {
	log.Println("正在解析并验证令牌")
	// 解析JWT令牌
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 验证令牌是否有效
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// 获取用户信息
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	/*
		根据JWT中的payload获取用户结构
	*/
	source := claims.Subject
	log.Println("令牌验证完毕")
	user := ToUser(source)
	return user, nil
}
