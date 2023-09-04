package client

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"simple/api-signature/client/utils"
	"strconv"
	"time"
)

/** 生成包含N个随机数字的字符串
 */
func GenetateRandomString(length int) string {
	// 设置随机数种子，以确保每次运行生成的随机数都不同
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 定义一个包含数字字符的字符集
	charset := "0123456789"
	charsetLength := len(charset)

	// 生成随机数字并拼接字符串
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(charsetLength)
		randomChar := charset[randomIndex]
		randomString[i] = randomChar
	}
	return string(randomString)
}

// 获得请求头
func GetRequestHeaders(accessKey, secretkey, requestBody string) http.Header {
	headers := make(http.Header)

	// 生成 nonce : 一个包含100个随机数字的字符串
	nonce := GenetateRandomString(100)

	// 当前时间戳（秒级别）
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 计算签名
	signature := utils.CalculateSignature(accessKey, secretkey, nonce, timestamp, requestBody)

	// 设置请求头
	headers.Set("accessKey", accessKey)
	headers.Set("nonce", nonce)
	headers.Set("timestamp", timestamp)
	headers.Set("sign", signature)

	return headers
}

func SendApi(name, accessKey, secretKey string) (statusCode int, contentType string, bodyBytes []byte, err error) {
	requestURL := "http://localhost:8011/api/name"

	// 构建查询字符串，将其附加到URL上
	params := url.Values{}
	params.Set("name", name)

	// 构建包含查询参数的URL
	fullURL := fmt.Sprintf("%s?%s", requestURL, params.Encode())

	client := &http.Client{}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Failed to create request, err=", err)
		return
	}

	// 构建请求头
	headers := GetRequestHeaders(accessKey, secretKey, "")
	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to make request, err=", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response, err=", err)
		return
	}

	statusCode = response.StatusCode
	contentType = response.Header.Get("Content-Type")

	return
}
