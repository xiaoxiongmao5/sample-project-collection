package main

import (
	"net/http"
	"simple/api-signature/service/utils"

	"github.com/gin-gonic/gin"
)

var AccessKey, SecretKey string = "aaa", "123456"

func GetNameByGet(c *gin.Context) {
	name := c.Query("name")

	headers := c.Request.Header
	accessKey := headers.Get("accessKey")
	nonce := headers.Get("nonce")
	timestamp := headers.Get("timestamp")
	sign := headers.Get("sign")

	if accessKey != AccessKey {
		c.JSON(http.StatusForbidden, gin.H{"error": "用户不存在"})
		return
	}

	// 计算签名
	signature := utils.CalculateSignature(accessKey, SecretKey, nonce, timestamp, "")

	// 验证签名
	if signature != sign {
		c.JSON(http.StatusForbidden, gin.H{"error": "签名验证失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "GET 你的名字是" + name})
}

func main() {
	r := gin.New()
	r.GET("/api/name", GetNameByGet)
	r.Run(":8011")
}
