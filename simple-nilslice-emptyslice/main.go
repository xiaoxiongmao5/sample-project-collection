package main

import (
	"fmt"
	"reflect"
)

// 测试 new 和 make 的地址分配问题
func main() {
	slice1 := make([]int, 0)
	slice2 := []int{}
	// 上面两种方式是一样的效果
	var slice3 []int
	slice4 := new([]int)
	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}
	if slice2 == nil {
		fmt.Println("slice2 is nil")
	}
	if slice3 == nil {
		fmt.Println("slice3 is nil")
	}
	if *slice4 == nil {
		fmt.Println("*slice4 is nil")
	}
	if reflect.DeepEqual(slice1, slice3) {
		fmt.Println("slice1 slice3相等")
	} else {
		fmt.Println("slice1 slice3不等")
	}

	/** 输出
	slice3 is nil
	*slice4 is nil
	slice1 slice3不等
	*/

	// slice1 = append(slice1, 1, 2, 3)
	// slice2 = append(slice2, 1, 2, 3)
	// slice3 = append(slice3, 1, 2, 3)
	// *slice4 = append(*slice4, 1, 2, 3)

	// if reflect.DeepEqual(slice1, slice2) {
	// 	fmt.Println("slice1 slice2相等")
	// } else {
	// 	fmt.Println("slice1 slice2不等")
	// }
	// if reflect.DeepEqual(slice1, slice3) {
	// 	fmt.Println("slice1 slice3相等")
	// } else {
	// 	fmt.Println("slice1 slice3不等")
	// }
	// if reflect.DeepEqual(slice3, slice4) {
	// 	fmt.Println("slice4 slice3相等")
	// } else {
	// 	fmt.Println("slice4 slice3不等")
	// }

	fmt.Printf("slice1 原值v=%v, 类型T=%T, 地址p=%p \n", slice1, slice1, &slice1)
	fmt.Printf("slice2 原值v=%v, 类型T=%T, 地址p=%p \n", slice2, slice2, &slice2)
	fmt.Printf("slice3 原值v=%v, 类型T=%T, 地址p=%p \n", slice3, slice3, &slice3)
	fmt.Printf("slice4 原值v=%v, 类型T=%T, 地址p=%p", slice4, slice4, &slice4)

	/** 输出
	slice1 原值v=[], 类型T=[]int, 地址p=0x14000116018
	slice2 原值v=[], 类型T=[]int, 地址p=0x14000116030
	slice3 原值v=[], 类型T=[]int, 地址p=0x14000116048
	slice4 原值v=&[], 类型T=*[]int, 地址p=0x14000122018%
	*/
}
