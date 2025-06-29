package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	// 设置Claims，Username可以由应用自由定义，ExpiresAt为过期时间（此为72小时）
	claims := CustomClaims{
		Username: "carpe",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 有效时间
			Issuer:    "wh",                                               // 签发人
			//IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			//NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			//Subject:   "somebody",                                         // 主题
			//ID:        "1",                                                // JWT ID用于标识jwt
			//Audience:  []string{"somebody_else"},                          // 用户

		},
	}

	// 创建时指定签名方式为 HMAC-SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用Secret进行签名，生产时需要设置好 JWT_SECRET 环境变量
	signed, err := token.SignedString([]byte("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	// 一般我们都会在Header中携带时需要 "Bearer " 前缀
	return "Bearer " + signed, nil
}

func main() {
	token, err := GenerateJWT("carpe")
	if err != nil {
		fmt.Println("生成jwt失败", err)
	}
	fmt.Println("token", token)

	//claims, err := jwt.ParseJWT("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNhcnBlIiwiZXhwIjoxNzUxNDM2NDMyfQ.NfeEqiMLYUaVHG5VSyk6nIb-R_x_dfNiSBdpIX5q6oA")
	//if err != nil {
	//	fmt.Println("ParseJWT失败", err)
	//}
	//fmt.Println("claims", claims)
}
