package Handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"testing"
	"yiliao/Dao"
	"yiliao/Utils"
)

var jwtKey = []byte("your-secret-key")

func TestName(t *testing.T) {
	authHeader := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUyNjU0NTgsImlzcyI6InlvdXItaXNzdWVyIiwic3ViIjoiYWRtaW4ifQ.g4oNat_FRTwbRWiDNJ0InMjybZNTunS5fF2qXxearHE"
	parts := strings.Split(authHeader, " ")
	// 获取第二部分（"Bearer" 之后的字符串）
	user, _ := Utils.ValidateToken(parts[1])

	fmt.Println("get user! ", user.Username)
}

func parsertoken(tokenstring string) *Dao.User {
	token, _ := jwt.ParseWithClaims(tokenstring, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	print("HEWEWLOWR")

	// 获取用户信息
	claims, _ := token.Claims.(*jwt.StandardClaims)
	print("HEWEWLOWR")

	fmt.Println("claims is :", claims)
	// 根据令牌中的用户信息构建用户结构
	user := &Dao.User{
		Username: claims.Subject,
	}
	return user
}
