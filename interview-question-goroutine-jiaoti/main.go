package main

import (
	"fmt"
	"sync"
)

// 题目：定义两个变量，a是全字母，b是全数字，定义2个线程分别打印a和b，想要2个线程能交替输出，先输出a 输出1，在输出b 再输出2，交替打印怎么去实现？
func main() {
	a := "abcdefghijklmn"
	b := "123456789"
	lena, lenb := len(a), len(b)
	length := max(lena, lenb)
	gp := sync.WaitGroup{}
	gp.Add(2)

	ch := make(chan int)
	go echo(b, ch, &gp, length)
	go echo(a, ch, &gp, length)

	ch <- 1
	gp.Wait()
	fmt.Println("ok")
}

func echo(str string, ch chan int, gp *sync.WaitGroup, length int) {
	arr := []byte(str)
	leng := len(arr)
	for i := 0; i < length; i++ {
		v := <-ch
		if i < leng {
			fmt.Println(string(arr[i]))
		}
		if i == length-1 && v%2 == 0 {
			fmt.Println("是最后一个，就不写ch了")
		} else {
			ch <- v + 1
		}
	}
	gp.Done()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
