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
	json.Unmarshal([]byte(str), &user)
	return &user
}

/*
创建JWT令牌
*/
func CreateToken(user Dao.User) string {
	// 设置令牌的过期时间
	expirationTime := time.Now().Add(24 * time.Hour)

	/*
		JWT令牌中存放用户对象
		将用户对象封装成Json字符串
		使用Base64Url编码进行加密
	*/
	userObj := ToJson(user)
	log.Println("userObj is :", userObj)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "your-issuer", // 可以自定义
		Subject:   userObj,
	}
	fmt.Println("claims is :", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名JWT令牌
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println("create Token :", tokenString)
	return tokenString
}

// ValidateToken 验证JWT令牌  ,参数为String 返回值为User
func ValidateToken(tokenString string) (*Dao.User, error) {
	fmt.Println("got token :", tokenString)
	// 解析JWT令牌
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 验证令牌是否有效
	if !token.Valid {
		log.Println(token)
		return nil, fmt.Errorf("invalid token")
	}

	// 获取用户信息
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	fmt.Println("claims is :", claims)
	/*
		根据JWT中的payload获取用户结构
	*/
	source := claims.Subject
	user := ToUser(source)
	return user, nil
}

func ConvertBoolToInt(flag bool) int {
	if flag {
		return 1
	} else {
		return 0
	}
}
