package jt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	Project = "bingo" //项目名称
)

var SecretKey = []byte("123456") //私钥

type JWTClaims struct {
	UID int64 `json:"uid"`
	jwt.StandardClaims
}

// NewJWT jwt分3部分组成，header,payload,signature
func NewJWT() *JWTClaims {
	maxAge := 2 // 过期时间,单位秒
	claims := &JWTClaims{
		// StandardClaims符合JWT标准的信息，比如nbf,exp等等
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    Project,                                                    // 非必须，也可以填充用户名，
		},
	}
	return claims
}

func CreateToken(uid int64) (string, error) {
	claims := NewJWT()
	// 假设只携带这一个自定义参数
	claims.UID = uid
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 查看过源码，必须是[]byte
	return token.SignedString(SecretKey)
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		// 该方法用来判断加密算法，返回私钥
		return SecretKey, nil
	})

	// 断言token是否有效
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
