package main

import (
	"context"
	"fmt"
	"time"
)

// golang context 上下文管理 context.Background() 介绍：https://www.cnblogs.com/wordgao/p/15735147.html

// 演示 context.Value() 的使用

func step1(ctx context.Context) context.Context {
	// 通过 WithValue 创建一个context
	child := context.WithValue(ctx, "name", "小花")
	return child
}
func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 12)
	return child
}
func step3(ctx context.Context) {
	fmt.Printf("name = %s \n", ctx.Value("name"))
	fmt.Printf("age = %d \n", ctx.Value("age"))
}

// 演示 context.Done() 的使用
func f1() {
	// 通过 WithTimeout 函数创建一个context
	//设置100毫秒的超时时间，到时会调用Deadline函数，就会关闭这个管道，Err() 里面就会保留这个管道为什么会关闭，信息回保存在error里面
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*100)
	defer cancel()
	select {
	// 读取管道数据（因为没有往里面写数据，所以管道是空的，所以读这里会阻塞，直到管道被关闭，才会结束阻塞）
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
	}
}

func f2() {
	parent, cancel1 := context.WithTimeout(context.TODO(), time.Millisecond*1000)
	defer cancel1()
	t0 := time.Now()

	time.Sleep(500 * time.Millisecond) //经过这里时，parent寿命还剩500ms

	child, cancel2 := context.WithTimeout(parent, time.Millisecond*1000) //这里设置child的寿命是1000ms，但其父ctx是500ms，所以这里以短的为准(不论父子，只论长短，以短的为准)
	defer cancel2()
	t1 := time.Now()

	select {
	case <-child.Done():
		err := child.Err()
		t3 := time.Now()
		fmt.Println("t3-t0", t3.Sub(t0).Milliseconds())
		fmt.Println("t3-t1", t3.Sub(t1).Milliseconds())
		fmt.Println(err)
	}
}

// 演示 cancel 的使用
func f3() {
	// 显示的调用cancel, 触发Deadline函数，就会关闭这个管道，Err() 里面就会保留这个管道为什么会关闭，信息回保存在error里面
	ctx, cancel := context.WithCancel(context.TODO())
	t0 := time.Now()
	// 通过一个子协程来调用cancel函数
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	select {
	// 读取管道数据（因为没有往里面写数据，所以管道是空的，所以读这里会阻塞，直到管道被关闭，才会结束阻塞）
	case <-ctx.Done():
		t3 := time.Now()
		fmt.Println("t3-t0", t3.Sub(t0).Milliseconds())
		err := ctx.Err()
		fmt.Println(err)
	}
}

func main() {
	// 演示 context.Value() 的使用
	// grandpa := context.TODO()
	// father := step1(grandpa)
	// grandson := step2(father)
	// step3(grandson)
	// 输出 >
	// name = 小花
	// age = 12

	// 演示 context.Done() 的使用
	// f1()
	// 输出 > context deadline exceeded

	// f2()
	// 输出 >
	// t3-t0 1001
	// t3-t1 499
	// context deadline exceeded

	// 演示 cancel() 的使用
	f3()
	// 输出 >
	// t3-t0 101
	// context canceled
}
