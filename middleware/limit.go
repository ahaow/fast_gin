package middleware

import (
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"time"
)

type Limiter struct {
	limit      int                // 限制的请求数量
	duration   time.Duration      // 时间窗口
	timestamps map[string][]int64 // 请求的时间戳
}

func LimitMiddleware(limit int) gin.HandlerFunc {
	return NewLimiter(limit, 1*time.Second).Middleware
}

var (
	limiter = NewLimiter(10, 1*time.Minute) // 每分钟最多请求 10 次
)

func NewLimiter(limit int, duration time.Duration) *Limiter {
	return &Limiter{
		limit:      limit,
		duration:   duration,
		timestamps: make(map[string][]int64),
	}
}

func (l *Limiter) Middleware(c *gin.Context) {
	ip := c.ClientIP() // 获取客户端ip地址

	// 检查请求时间戳切片是否存在
	if _, ok := l.timestamps[ip]; !ok {
		l.timestamps[ip] = make([]int64, 0)
	}

	now := time.Now().Unix() // 获取当前时间戳

	// 移除过期时间戳
	for i := 0; i < len(l.timestamps[ip]); i++ {
		if l.timestamps[ip][i] < now-int64(l.duration.Seconds()) {
			l.timestamps[ip] = append(l.timestamps[ip][:i], l.timestamps[ip][i+1:]...)
			i--
		}
	}

	// 检查请求数量是否超过限制
	if len(l.timestamps[ip]) >= l.limit {
		res.FailWithMsg("Too Many Request", c)
		//c.JSON(429, gin.H{
		//	"code": 500,
		//	"data": gin.H{},
		//	"msg":  "Too Many Request",
		//})
		c.Abort()
		return
	}

	// 添加当前请求时间戳
	l.timestamps[ip] = append(l.timestamps[ip], now)

	// 继续处理请求
	c.Next()
}
