package main

// ./kafka-console-consumer.sh --topic user_click --bootstrap-server localhost:9092 --from-beginning --group a
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	topic  = "user_click"
)

func read(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		if message, err := reader.ReadMessage(ctx); err != nil {
			fmt.Println("读kafka失败: ", err)
			break
		} else {
			fmt.Printf("topic=%s, partition=%d, offset=%d, key=%s, value=%s \n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		}
	}
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Println("接收到信号: ", sig.String())
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()
	go listenSignal()
	read(ctx)
	// to consume messages
	// topic := "my-topic"
	// partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }

	// conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	// b := make([]byte, 10e3) // 10KB max per message
	// for {
	// 	n, err := batch.Read(b)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(b[:n]))
	// }

	// if err := batch.Close(); err != nil {
	// 	log.Fatal("failed to close batch:", err)
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close connection:", err)
	// }
}
