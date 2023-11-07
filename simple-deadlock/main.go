package main

import (
	"fmt"
	"sync"
	"time"
)

/**
死锁的几种场景代码
*/

func main() {
	DeadLock5()
}

// 只写不读，主协程通过 sleep 阻塞
func DeadLock() {
	ch := make(chan int, 0)
	go func() {
		ch <- 1
		fmt.Println("over")
	}()
	// 这里不会马上返回deadlock，会进入休眠，因为休眠本身时间到了之后会自行解锁。而协程并不知道后面是否有解除主协程的代码或者解除协程1的代码
	time.Sleep(3 * time.Second)
	// 无输出
}

// 只写不读，主协程通过 写管道 阻塞
func DeadLock2() {
	ch := make(chan int, 0)
	go func() {
		ch <- 1
		fmt.Println("over")
	}()
	ch <- 1
	// 输出 fatal error: all goroutines are asleep - deadlock!
}

// 只写不读，主协程通过 WaitGroup 阻塞
func DeadLock3() {
	ch := make(chan int, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		fmt.Println("over")
	}()
	wg.Wait()
	// 输出 fatal error: all goroutines are asleep - deadlock!
}

// 只写不读，主协程通过 WaitGroup 阻塞，新增一个子协程通过 sleep 阻塞
func DeadLock4() {
	ch := make(chan int, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		fmt.Println("over")
	}()
	go func() {
		// 这里不会马上返回deadlock，会进入休眠，因为休眠本身时间到了之后会自行解锁。而协程并不知道后面是否有解除主协程的代码或者解除协程1的代码
		time.Sleep(3 * time.Second)
		// wg.Done()
		// <-ch
	}()
	wg.Wait()
	// 输出 fatal error: all goroutines are asleep - deadlock!
}

// 只写不读
func DeadLock5() {
	ch := make(chan int, 0)
	ch <- 1
	fmt.Println("ok")
	// 输出 fatal error: all goroutines are asleep - deadlock!
}
