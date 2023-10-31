package main

import "fmt"

func main() {
	ch1 := make(chan int, 1) //有缓存通道，可以写
	ch1 <- 1                 //可以写
	fmt.Println("ch1 有缓存通道，可以直接写")

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
