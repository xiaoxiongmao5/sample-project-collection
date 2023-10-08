// 服务器1
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Arith struct{}

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Args struct {
	A, B int
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Println("RPC server is running on port 1234...")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
