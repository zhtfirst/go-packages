package example

import (
	"fmt"

	"github.com/zhtfirst/go-packages/jwt"
)

func Jwt() {
	jwt.InitJWTToken("secret") // secret 为密钥

	// 生成令牌
	token, _ := jwt.JWTToken.CreateToken("张三", 20)
	fmt.Println("生成令牌为:", token)

	// 解析令牌
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjMyMjYzMTQsImlzcyI6IlRpbWVUb2tlbiIsInN1YiI6IuW8oOS4iSJ9.fZHj39IygfniacFi20crm9kHrvNIJJeZM91Fw89QVpM"
	claims, _ := jwt.JWTToken.ParseToken(tokenString)
	fmt.Println("解析参数为：", claims)

}
