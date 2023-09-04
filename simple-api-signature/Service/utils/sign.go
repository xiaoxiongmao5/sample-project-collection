package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算API签名
func CalculateSignature(accessKey, secretKey, nonce, timestamp, requestBody string) string {
	// 将参数拼接成一个字符串
	concatenatedString := accessKey + nonce + timestamp + requestBody + secretKey

	// 计算 MD5 值
	signature := md5.Sum([]byte(concatenatedString))
	return hex.EncodeToString(signature[:])
}
