package main

import "fmt"

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}

func main() {
	defer func() {
		fmt.Println("c")
	}()
	//子函数抛出的panic没有recover时，上层函数时，程序直接异常终止
	F()
	fmt.Println("继续执行")
}
