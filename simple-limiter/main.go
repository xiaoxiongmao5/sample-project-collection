package main

import (
	"net/http"
	"simple-limiter/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建一个限流器，每秒允许最多10个请求
	// middleware.Limiter = rate.NewLimiter(rate.Every(100*time.Millisecond), 1)
	// middleware.Limiter = rate.NewLimiter(rate.Limit(10), 1) //每隔100毫秒生产1个令牌

	// 使用限流中间件
	// r.Use(middleware.RateLimiterMiddleware())

	// 创建IP限流器
	middleware.IPLimiter = middleware.NewIPRateLimiter()

	// 使用限流中间件
	r.Use(middleware.IPRateLimiterMiddleware())

	// 定义一个路由处理函数
	r.GET("/api/resource", func(c *gin.Context) {
		time.Sleep(50 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{"message": "Resource accessed"})
	})

	// 启动Gin应用
	r.Run(":8080")
}
