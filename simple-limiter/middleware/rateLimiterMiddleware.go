package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var Limiter *rate.Limiter

// 定义一个中间件函数来进行限流
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !Limiter.AllowN(time.Now(), 1) {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			// 设置休眠和业务时长一样，为了更好从日志出看出规则
			time.Sleep(50 * time.Millisecond)
			c.Abort()
			return
		}

		c.Next()
	}
}
