package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySigningKey = []byte("secret")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
func ParseAndVerifyToken(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名算法异常：%v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Printf("token 验证失败，%v\n", err)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("token 验证通过")
		fmt.Printf("当前用户：%v\n", claims["sub"])
		fmt.Printf("过期时间 %v\n", time.Unix(int64(claims["exp"].(float64)), 0))
	} else {
		fmt.Println("token 无效")
	}
}
func main() {
	username := "student_zhang"
	tokenString, err := GenerateToken(username)
	if err != nil {
		return
	}
	fmt.Println(tokenString)
	ParseAndVerifyToken(tokenString)
	fmt.Println("模拟修改")
	fakeToken := tokenString[:20] + "hacked_data" + tokenString[20:]
	ParseAndVerifyToken(fakeToken)
}
