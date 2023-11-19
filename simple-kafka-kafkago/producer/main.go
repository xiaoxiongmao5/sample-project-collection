package main

// ./kafka-console-producer.sh --topic my-topic --bootstrap-server localhost:9092

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	topic  = "user_click"
)

func Write(ctx context.Context) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	tm := time.Now().Unix()
	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("小" + strconv.Itoa(int(tm)))},
			kafka.Message{Key: []byte("2"), Value: []byte("白" + strconv.Itoa(int(tm)))},
			kafka.Message{Key: []byte("3"), Value: []byte("小" + strconv.Itoa(int(tm)))},
			kafka.Message{Key: []byte("1"), Value: []byte("熊" + strconv.Itoa(int(tm)))},
			kafka.Message{Key: []byte("1"), Value: []byte("猫" + strconv.Itoa(int(tm)))},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Println("批量写kafka失败：%v \n", err)
			}
		} else {
			break
		}
	}
}

func main() {
	ctx := context.Background()
	Write(ctx)
	// to produce messages
	// topic := "my-topic"
	// partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	// if err != nil {
	// 	log.Fatal("failed to write messages:", err)
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// }
}
