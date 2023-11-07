package main

import (
	"fmt"
	"sync"
	"time"
)

/**
管道-通道-Channel
0. 基本概念：管道底层是一个【环形队列】，先进先出。特别适合多协程并发的场景。
	* 一旦在多协程情况下，需要多个协程共同去读写某个容器，这种情况应该先想到管道。
	* make(chan int, 8)		8 代表容量，不是长度，即刚创建里面还是没有元素的。
1. 将管道应用于一个生产者-消费者模式当中
2. 管道的阻塞问题
	* sync.WaitGroup{}
3. 如何通过管道实现多个协程的协调与同步
	* make(chan struct{}, 0)
*/

// 1. 将管道应用于一个生产者-消费者模式当中
func ProducerConsumer() {
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 2个生产者,往channel内写元素
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// wg2 := sync.WaitGroup{}
	// wg2.Add(1)
	mc := make(chan struct{}, 0) //空结构体，不占任何内存空间，并且增强可读性（该管道不是用做容器使用，而是充当用于一个协程之间的协调同步作用）
	// 1个消费者
	go func() {
		// defer wg2.Done()
		defer func() {
			mc <- struct{}{}
		}()
		sum := 0
		for {
			v, ok := <-ch
			if !ok { //channel被关闭，且channel为空，此时OK才为false
				break
			}
			sum += v
		}
		fmt.Println("sum=", sum)
	}()

	wg.Wait()
	close(ch) //意味着，从此之后管道不可用了，不允许再往管道里面写入元素了，但仍然可以读取
	// wg2.Wait()
	<-mc
}

func main() {
	ProducerConsumer()

	ch := make(chan int, 0) //容量为0的管道
	fmt.Println(time.Now().Unix())
	go func() {
		ch <- 4
		fmt.Println(time.Now().Unix(), "写入4到管道成功")
	}()
	time.Sleep(2 * time.Second)
	go func() {
		v := <-ch
		fmt.Println(time.Now().Unix(), "读取管道成功, v=", v)
	}()
	time.Sleep(1 * time.Second)
	/** 上面代码的输出结果：
	1699344916
	1699344918 写入4到管道成功
	1699344918 读取管道成功, v= 4
	*/

	ch02 := make(chan int, 1) //容量为1的管道
	fmt.Println(time.Now().Unix())
	go func() {
		ch02 <- 5
		fmt.Println(time.Now().Unix(), "写入5到管道成功")
	}()
	time.Sleep(2 * time.Second)
	go func() {
		v := <-ch02
		fmt.Println(time.Now().Unix(), "读取管道成功, v=", v)
	}()
	time.Sleep(1 * time.Second)
	/** 上面代码的输出结果：
	1699345176
	1699345176 写入5到管道成功
	1699345178 读取管道成功, v= 5
	*/

	ch03 := make(chan int, 1)
	fmt.Println(time.Now().Unix())
	go func() {
		v := <-ch03
		fmt.Println(time.Now().Unix(), "读取管道成功, v=", v)
	}()
	time.Sleep(2 * time.Second)
	go func() {
		ch03 <- 6
		fmt.Println(time.Now().Unix(), "写入6到管道成功")
	}()
	time.Sleep(1 * time.Second)
	/** 上面代码的输出结果：
	1699345458
	1699345460 写入6到管道成功
	1699345460 读取管道成功, v= 6
	*/

	// ch1 := make(chan int, 1) //有缓存通道，可以写
	// ch1 <- 1                 //可以写
	// fmt.Println("ch1 有缓存通道，可以直接写")

	// ch2 := make(chan int) //无缓存通道
	// ch2 <- 1              //fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("ch2 无缓存通道，不能直接写")

	// ch3 := make(chan int, 0) //有缓存通道长度=0  ===  无缓存通道
	// ch3 <- 1                 //fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("ch3 无缓存通道，不能直接写")

	// ch4 := make(chan int, 0) //有缓存通道长度=0  ===  无缓存通道
	// <-ch4                    //fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("ch4 无缓存通道，不能直接写")

	// ch5 := make(chan int) //有缓存通道长度=0  ===  无缓存通道
	// <-ch5                 //fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("ch5 无缓存通道，不能直接写")
}

/**
总结：

定义一个有缓存通道，容量=0，跟无缓存通道有什么区别？能写吗，能读吗？

`ch := make(chan int, 0) ` 这段代码创建的通道是无缓存通道（Unbuffered Channel），因为通道的容量被明确设置为0

当创建一个长度为0的通道时，它就成为了无缓存通道，而不管是否使用make(chan T)还是make(chan T, 0)来创建它，效果都是一样的。

在无缓存通道中，数据发送和接收操作是同步的，发送方会等待接收方，接收方会等待发送方。

这种通道用于强制同步，确保数据的安全传递，通常用于等待其他协程的响应或结果。它确保了数据的传输是即时的，而不会被缓冲。

*/
