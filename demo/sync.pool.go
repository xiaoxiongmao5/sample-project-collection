package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var createNum int32

// 创建一个长度为 1024 字节的字节切片，并原子的递增 createNum，用于计数创建的对象数量
func createBuffer() interface{} {
	atomic.AddInt32(&createNum, 1)
	buffer := make([]byte, 1024)
	return buffer
}

func main() {
	// 创建 Pool 对象，并指定用于创建新对象的函数 createBuffer
	bufferPool := &sync.Pool{New: createBuffer}

	workerPool := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerPool)

	for i := 0; i < workerPool; i++ {
		go func() {
			defer wg.Done()
			// 从对象池获取一个对象（如果对象池为空，则会调用 createBuffer 来创建新对象）
			buffer := bufferPool.Get()
			// 将获取的对象断言为 []byte 类型
			_ = buffer.([]byte)

			// buffer := createBuffer()
			// _ = buffer.([]byte)

			// 将对象放回对象池，以便重用
			defer bufferPool.Put(buffer)
		}()
	}
	// 等待所有 worker 协程执行完毕
	wg.Wait()
	fmt.Printf(" %d buffer objects were create.\n", createNum)
	time.Sleep(3 * time.Second)
}
