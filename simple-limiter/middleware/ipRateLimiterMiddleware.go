package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var IPLimiter *IPRateLimiter

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		limiter: make(map[string]*rate.Limiter),
	}
}

type IPRateLimiter struct {
	mu      sync.Mutex
	limiter map[string]*rate.Limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.limiter[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(2), 5) // 每秒2个请求，桶容量为5
		i.limiter[ip] = limiter
	}

	return limiter
}

// 定义一个中间件函数来进行限流
func IPRateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter := IPLimiter.GetLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			// 设置休眠和业务时长一样，为了更好从日志出看出规则
			time.Sleep(50 * time.Millisecond)
			c.Abort()
			return
		}

		c.Next()
	}
}
