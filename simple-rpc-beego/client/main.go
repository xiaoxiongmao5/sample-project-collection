// 客户端
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	// 客户端通过 rpc.Dial 连接到服务器
	// 连接到服务器1
	client1, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing server1:", err)
	}

	// 连接到服务器2
	client2, err := rpc.Dial("tcp", "localhost:5678")
	if err != nil {
		log.Fatal("dialing server2:", err)
	}

	args := &Args{7, 8}
	var reply1, reply2 int

	// 使用 client.Call 来调用服务器上的 Arith.Multiply 方法。
	// 在服务器1上调用 Multiply 方法
	err = client1.Call("Arith.Multiply", args, &reply1)
	if err != nil {
		log.Fatal("arith error on server1:", err)
	}

	// 在服务器2上调用 Multiply 方法
	err = client2.Call("Arith.Multiply", args, &reply2)
	if err != nil {
		log.Fatal("arith error on server2:", err)
	}

	fmt.Printf("Arith on server1: %d*%d=%d\n", args.A, args.B, reply1)
	fmt.Printf("Arith on server2: %d*%d=%d\n", args.A, args.B, reply2)
}
