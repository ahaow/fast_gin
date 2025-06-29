package middleware

import (
	"fast_gin/utils/res"

	"github.com/gin-gonic/gin"
)

// ✅ 泛型 JSON 绑定函数（在 handler 中用）
func BindJSON[T any](c *gin.Context) (T, bool) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return req, false
	}
	return req, true
}

// ✅ 泛型 Query 参数绑定
func BindQuery[T any](c *gin.Context) (T, bool) {
	var req T
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return req, false
	}
	return req, true
}

// ✅ 泛型 URI 参数绑定（如 path 中的 id）
func BindUri[T any](c *gin.Context) (T, bool) {
	var req T
	if err := c.ShouldBindUri(&req); err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return req, false
	}
	return req, true
}

// ✅ 泛型 JSON 中间件（用于中间件链）
func BindJsonMiddleware[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			res.FailWithError(err, c)
			c.Abort()
			return
		}
		c.Set("bind_request", req)
		c.Next()
	}
}

// ✅ 提取绑定数据（中间件方式时使用）
func GetBind[T any](c *gin.Context) T {
	return c.MustGet("bind_request").(T)
}
