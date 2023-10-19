/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 11:35:24
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 11:55:52
 * @FilePath: /simple-redis-msgqueue/consumer/consumer.go
 * @Description: 消费者
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

	// 启动消费者协程
	go consumeMessages(ctx, client)

	// 主进程可以继续执行其他操作
	fmt.Println("主进程继续执行...")

	// 主进程可以等待一段时间或执行其他操作
	time.Sleep(10 * time.Minute) // 例如，等待10分钟
}

func consumeMessages(ctx context.Context, client *redis.Client) {
	for {
		select {
		case <-ctx.Done():
			return // 如果上下文被取消，则退出协程
		default:
			// 从队列的左侧（头部）获取消息，这将阻塞等待直到有消息可用
			message, err := client.LPop(ctx, "myqueue").Result()
			if err == redis.Nil {
				// 队列为空，等待一段时间后继续尝试
				time.Sleep(time.Second)
				continue
			} else if err != nil {
				fmt.Println("无法获取消息:", err)
				return
			}

			fmt.Printf("已消费消息: %s\n", message)

			// 在这里添加处理消息的逻辑，例如执行某些操作

			// 模拟处理消息所需的时间
			time.Sleep(time.Second)
		}
	}
}
