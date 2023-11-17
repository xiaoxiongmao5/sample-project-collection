package jsontest

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

/**
单元测试：
	* 测试文件名称：_test.go 结尾
	* 函数名称：Test 作为前缀，参数必须是 *testing.T
	* -v 是单元测试函数内的 print 正常输出
	* -run 正则表达式(不需要写全函数名也可以)，明确要运行的函数
 	* -count=1 禁止测试缓存（1 表示运行测试的次数，如果是2代表连续运行2次） （go测试工具默认会在运行测试前移除测试二进制文件(test binary)的缓存。这是为了确保每次运行测试时都能使用最新的代码。如果想禁用缓存，可以使用`-count`标志，将其设置为非零值。这会告诉测试工具在每次运行测试时都重新编译并运行。）
	* 运行命令比如：go test -v -count=1 ./json_test -run=Sonic

基准测试：(就是性能测试，反复调用很多次)
	* 注意：不关心逻辑是否正确，只关心性能
	* 函数名称：Benchmark 作为前缀，参数必须是 *testing.B
		* for循环，b.N 不是一个常量，会根据你实际的运行耗时来动态的选择 b.N，总之是个很大的数
	* -bench 指定测试哪个函数
	* -benchmem 查看内存使用和申请情况
	* 运行命令比如：go test -bench=Json json_test.go -benchmem
	* 输出解析：
		goos: darwin【操作系统】
		goarch: arm64【CPU架构】
		cpu:【CPU型号】
		BenchmarkUseJson-8	492722【一共运行了多少次，就是N的值】	2414 ns/op【每次运行消耗的时间】	736 B/op【每次操作大概消耗这么多内存】	19 allocs/op【每次操作大概进行了19次的内存申请】

				补充：	operation：一个操作 => /op 每次操作

*/

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s = Student{"小花", 18, true}
	c = Class{"1(2)班", []Student{s, s, s}}
)

// 标准库比sonic慢了 3、4 倍
// func TestUseJson(t *testing.T) {
// 	bytearr, err := json.Marshal(c)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	var c2 Class
// 	err = json.Unmarshal(bytearr, &c2)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if !(c2.Id == c.Id && len(c2.Students) == len(c.Students)) {
// 		t.Fail()
// 	}
// 	fmt.Println("测试成功 by json")
// }

// func TestUseSonic(t *testing.T) {
// 	bytearr, err := sonic.Marshal(c)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	var c2 Class
// 	err = sonic.Unmarshal(bytearr, &c2)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if !(c2.Id == c.Id && len(c2.Students) == len(c.Students)) {
// 		t.Fail()
// 	}
// 	fmt.Println("测试成功 by sonic")
// }

func BenchmarkUseJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytearr, _ := json.Marshal(c)
		var c2 Class
		json.Unmarshal(bytearr, &c2)
	}
}

func BenchmarkUseSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytearr, _ := sonic.Marshal(c)
		var c2 Class
		sonic.Unmarshal(bytearr, &c2)
	}
}
