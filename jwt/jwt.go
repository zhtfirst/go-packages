package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func InitJWTToken(key string) {
	JWTToken = NewJWT(key)
}

var hmacSampleSecret []byte

type JWT struct {
	SigningKey []byte // 密钥
}

var JWTToken *JWT

func NewJWT(secretKey string) *JWT {
	return &JWT{
		[]byte(secretKey),
	}
}

// CreateToken 创建一个token
// param: data 数据
// param: timeout 过期时间（秒）
func (j *JWT) CreateToken(data string, timeout int64) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:  "TimeToken",
		Subject: data,
	}
	if timeout > 0 {
		claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeout)).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(token string) (string, error) {
	var err error
	var claims jwt.StandardClaims
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	err = claims.Valid()
	if err != nil {
		return "", err
	} else {
		return claims.Subject, err
	}
}
