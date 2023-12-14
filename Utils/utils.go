package Utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"yiliao/Dao"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("4295d277e69c8be76df37abbae2d989e9d8482dd")

/*
将用户对象转换为Json对象
*/
func ToJson(user Dao.User) string {
	jsonData, _ := json.Marshal(user)
	return string(jsonData)
}

/*
将Json对象转换为User对象
*/
func ToUser(str string) *Dao.User {
	var user Dao.User
	json.Unmarshal([]byte(str), &user)
	return &user
}

/*
使用Base64加密字符串
*/
func Base64UrlEncode(data string) string {
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	return fmt.Sprintf("%s", encodedData)
}

/*
使用Base64Url解密字符串
*/
func Base64UrlDecode(str string) string {
	decodedData, _ := base64.URLEncoding.DecodeString(str)
	return string(decodedData)
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

// 验证JWT令牌  ,参数为String 返回值为User
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
		print(token)
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
	jstr := Base64UrlDecode(source)
	user := ToUser(jstr)
	return user, nil
}

func ConvertBoolToInt(flag bool) int {
	if flag {
		return 1
	} else {
		return 0
	}
}
