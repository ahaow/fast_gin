package middleware

import (
	"fast_gin/utils/jwt"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	_, err := jwt.ParseJWT(token)
	fmt.Println("AuthMiddleware", err)
	if err != nil {
		// token失效
		res.FailWithMsg("认证失败", c)
		c.Abort()
		return
	}
	c.Next()
}
