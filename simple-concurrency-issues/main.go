package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
并发安全性 的解决方案：
1. 排他互斥锁 sync.Mutex、读写互斥锁 sync.RWMutex
	* sync.Mutex：用于在同一时刻只允许一个协程访问共享资源。
		* 当一个协程获得了锁（调用 Lock 方法），其他协程就会被阻塞，直到持有锁的协程释放它（调用 UnLock 方法）。
		* 适用于：对于临界区进行写操作，确保写操作的互斥性，防止数据竞争。
		* 场景：适用于需要排他性写访问的场景。

	* sync.RWMutex：允许多个协程同时访问共享资源，但在写操作时会互斥。
		* 当一个协程获得读锁（调用 RLock 方法），其他协程可以继续获取读锁，而写锁会阻塞其他协程。
		* 当一个协程获得写锁（调用 Lock 方法），其他协程无法获取读锁或写锁，直到写锁被释放。
		* 适用于：对于临界区进行读操作，允许多个协程同时读取，但在写操作时需要互斥，以确保写操作的互斥性。
		* 场景：适用于【读多写少】的场景，允许多个协程同时读取共享资源，但是写操作时保证互斥性。

2. 原子操作 sync/atomic
*/

var (
	a     int32
	lock  sync.Mutex   //互斥锁
	lock2 sync.RWMutex //读写互斥锁
)

// 方法1：互斥锁，在具体操作前上锁，处理完解锁
func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		a++
		lock.Unlock()
	}
}

func fn1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock2.Lock()
		a++
		lock2.Unlock()
	}
}

// 方法2：使用golang的原子操作
func fn2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&a, 1)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go fn2(&wg)
	go fn2(&wg)
	wg.Wait()
	fmt.Println("a=", a)
}
