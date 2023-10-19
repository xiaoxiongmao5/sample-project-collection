/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 11:33:51
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 11:52:51
 * @FilePath: /simple-redis-msgqueue/producer/producer.go
 * @Description: 生产者
 */
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端连接
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
		DB:   0,
	})

	defer client.Close()

	ctx := context.Background()

	// 生产者将消息推送到队列
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Message %d", i)

		// 将消息推送到队列的右侧（尾部）
		if err := client.RPush(ctx, "myqueue", message).Err(); err != nil {
			fmt.Println("无法将消息推送到队列:", err)
			return
		}

		fmt.Printf("已生产消息: %s\n", message)

		// 等待一段时间以模拟生产速度
		time.Sleep(time.Second)
	}
}
