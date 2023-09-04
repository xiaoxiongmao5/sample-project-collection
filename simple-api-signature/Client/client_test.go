package client

import (
	"fmt"
	"testing"
)

var AccessKey, SecretKey string = "aaa", "123456"

func TestSendApi(t *testing.T) {

	statusCode, contentType, bodyBytes, err := SendApi("xxx123", AccessKey, SecretKey)

	if err != nil {
		fmt.Println("ERROR: Failed to read response, err=", err)
		return
	}
	fmt.Printf("SUCCESS: statusCode=%v contentType=%v, bodyBytes=%v \n", statusCode, contentType, string(bodyBytes))
}
