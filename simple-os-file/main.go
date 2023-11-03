package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// go 读写文件

// 读文件
func ReadFile() {
	// 打开文件，不存在返回err
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 读文件
	bytearr := make([]byte, 100)
	num, err := file.Read(bytearr)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	fmt.Println("读取文件的字节数", num)
	fmt.Println("bytearr", string(bytearr[:num]))
	fmt.Println("bytearr", bytearr)
}

// 写文件
func WriteFile() {
	// 打开文件，不存在就创建 O_TRUNC-覆盖 O_APPEND-追加
	file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	var str = "小熊\na"
	num, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println("写文件失败", err)
		return
	}
	fmt.Printf("写入%d个字节的数据 \n", num)
}

// 带缓存的读取文件(默认一次读取4096字节=4M的数据)
func Do() {
	// 打开文件
	file, err := os.Open("b.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// bytearr, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("一次性读文件失败", err)
	// 	return
	// }
	// fmt.Print(string(bytearr))

	// fmt.Println()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF { // End Of File
				fmt.Print(line)
				break
			} else {
				fmt.Println("读文件发生错误", err)
				return
			}
		} else {
			fmt.Print(line)
		}
	}
}

func main() {
	// ReadFile()
	// WriteFile()
	Do()
}

/**
我们知道：
读写文件本身是很慢的操作，因为CPU每次跟磁盘进行交互的话，这个过程会很慢。
所以，为了减少CPU跟磁盘交互的次数，每次读写会读很多或写很多，把它暂时缓存起来，留给下一次使用。
这就是 buffer 的作用，主要是为了提高程序的性能。
通过一个普通的file，把它封装成一个带有缓冲的 reader

*/
