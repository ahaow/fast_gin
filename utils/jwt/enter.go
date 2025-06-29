package jwt

import (
	"errors"
	"fast_gin/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 一定要设置好 .env 中的 JWT_SECRET 值，生产时不可泄漏
// var Secret = []byte(os.Getenv("JWT_SECRET"))

// CustomClaims 定义我们自己的Claims（可以放需要携带到JWT中的数据）
// jwt.RegisteredClaims 是 JWT 本身定义的一系列字段（包括过期时间）
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT 函数签发 JWT，payload 中包括Username和过期时间
// 这通常是在认证时（login时）由后端颁发给客户端
func GenerateJWT(username string) (string, error) {
	jwtConfig := global.Config.Jwt
	// 设置Claims，Username可以由应用自由定义，ExpiresAt为过期时间（此为72小时）
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), // 有效时间
			Issuer:    global.Config.Jwt.Issuer,                                                                 // 签发人
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
	signed, err := token.SignedString([]byte(jwtConfig.Key))
	if err != nil {
		return "", err
	}

	// 一般我们都会在Header中携带时需要 "Bearer " 前缀
	return "Bearer " + signed, nil
}

// ParseJWT 函数用来验证 JWT 的合法性，并解析其中的数据（Claims）
// 一般是在需要认证时使用，如中间件中验证每一个API的Authorization头
func ParseJWT(signed string) (*CustomClaims, error) {
	jwtConfig := global.Config.Jwt

	if len(signed) > 7 && signed[:7] == "Bearer " {
		signed = signed[7:]
	}

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(signed, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Key), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
