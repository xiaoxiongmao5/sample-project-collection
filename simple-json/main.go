package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

func UseJson() {
	start := time.Now()
	s := Student{"小花", 18, true}
	c := Class{"1(2)班", []Student{s, s, s}}

	bytearr, err := json.Marshal(c) //JSON 序列化
	if err != nil {
		fmt.Println("json序列化失败", err)
		return
	}
	fmt.Println(string(bytearr))

	var c2 Class
	err = json.Unmarshal(bytearr, &c2)
	if err != nil {
		fmt.Println("json 反序列化失败", err)
		return
	}
	fmt.Printf("%+v", c2)
	total := time.Since(start).Microseconds()
	fmt.Println("耗时（纳秒）", total) //430
}

// 字节跳动研发的JSON序列化反序列化工具，比Go官方的性能上快很多
func UseSonic() {
	start := time.Now()
	s := Student{"小花", 18, true}
	c := Class{"1(2)班", []Student{s, s, s}}

	bytearr, err := sonic.Marshal(c) //JSON 序列化
	if err != nil {
		fmt.Println("json序列化失败", err)
		return
	}
	fmt.Println(string(bytearr))

	var c2 Class
	err = sonic.Unmarshal(bytearr, &c2)
	if err != nil {
		fmt.Println("json 反序列化失败", err)
		return
	}
	fmt.Printf("%+v", c2)
	total := time.Since(start).Microseconds()
	fmt.Println("耗时（纳秒）", total) //27
}

func main() {
	UseJson()
	fmt.Println()
	UseSonic()
}

/**
JSON 序列化：把一个对象转为一个二进制流（说白了就是一个 []byte）。
因为只有把它转为二进制之后，才能执行 IO 的存储，或者是网络的传输。
因为写在磁盘里的或是网上传输的，本质上都是一个二进制流，而不是一个结构体。

JSON 反序列化：把一个JSON的二进制流转为一个结构体。
*/
