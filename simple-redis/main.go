package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

/**
用Go语言读写 Redis
* 缓存就是万金油
* 应用场景：一般适用于以下两个场景
*	1. 在微服务的场景下，有多台服务器，它们需要并发的操作同一个变量。
		比如 a 服务器修改了某一个变量，希望 b 服务器在读取的时候，读到的是已修改后的最新值。
		这种时候，显然该值不能只存在 a 服务器上，而应该存在公共的地方，这个公共的地方就是分布式缓存。
	2. 有时候，需要加载大量数据放入缓存，如果在每台服务器上都开辟一个很大的内存空间来存放这些数据，显然非常消耗内存。
		将这部分内存，挪到一个公共的地方去，这就是分布式缓存的应用场景。
* 用到第三方的库：github.com/redis/go-redis/v9
*/

func string(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "小小熊猫"
	err := client.Set(ctx, key, value, 1*time.Second).Err() //第四个参数过期时间：0 表示永久生效，永不过期
	checkErr(err)

	// 设置过期时间
	client.Expire(ctx, key, 3*time.Second)

	time.Sleep(2 * time.Second)

	v2, err := client.Get(ctx, key).Result()
	checkErr(err)

	fmt.Println("v2=", v2)

	// 删除key
	client.Del(ctx, key)
}

func list(ctx context.Context, client *redis.Client) {
	key := "idx"
	values := []interface{}{1, 2, 3, "中国", "小熊"}
	// 如果该key不存在，会自动创建
	err := client.RPush(ctx, key, values...).Err()
	checkErr(err)

	// 读取区间（双闭区间）
	v2, err := client.LRange(ctx, key, 0, -1).Result() //0 第一个元素， -1 最后一个元素
	checkErr(err)
	fmt.Println("v2=", v2)

	client.Del(ctx, key)
}

func hashtable(ctx context.Context, client *redis.Client) {
	err := client.HSet(ctx, "学生1", "Name", "张三", "Age", 18, "Height", 175.8).Err()
	checkErr(err)
	err = client.HSet(ctx, "学生2", "Name", "李四", "Age", 20, "Height", 180.0).Err()
	checkErr(err)

	// 指定 field 读取
	v1, err := client.HGet(ctx, "学生1", "Name").Result()
	checkErr(err)
	fmt.Println("学生1的Name=", v1)
	v2, err := client.HGet(ctx, "学生2", "Age").Result()
	checkErr(err)
	fmt.Println("学生2的Age=", v2)

	// 指定 Key 读取
	for field, value := range client.HGetAll(ctx, "学生1").Val() {
		fmt.Println(field, value)
	}

	client.Del(ctx, "学生1")
	client.Del(ctx, "学生2")
}

func main() {
	// 创建链接
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, //启动后，默认创建0-15（共16个）DB，默认是0号DB
	})
	ctx := context.TODO()
	// string(ctx, client)
	// list(ctx, client)
	hashtable(ctx, client)
}

func checkErr(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
			return
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
