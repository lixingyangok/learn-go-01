package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1维数组
	// 数组1：固定长度3，整型，内容：空
	// 数组2：固定长度4，字符型，内容：A, B, C
	// 数组3：自动长度x，布尔型，内容：true, true
	// 数组4：固定长度3，字符型，内容：首位A，二位空，末位C
	var arr01 = [3]int{} //必得有花括号结尾
	var arr02 = [4]string{"A","B","C","D"}
	var arr03 = [...]bool{true, false}
	var arr04 = [3]string{0:"A", 2:"C"}
	fmt.Println(reflect.TypeOf(arr01), arr01)
	fmt.Println(reflect.TypeOf(arr02), arr02)
	fmt.Println(arr03)
	fmt.Println(arr04)

	
	// 2维数组
	// 数组1：长度为2的数组，成员也是2位长度数组，内容为字符
	var arr11 = [2][2]string{
		[2]string{"0-0", "0-1"},
		[2]string{"1-0", "1-1"},
	}
	fmt.Println(reflect.TypeOf(arr11), arr11)

	// ▼遍历
	for idx, val := range arr02 { 
		fmt.Println(idx, "--", val)
	}
	fmt.Println("-------------------------")
	for idx, val := range arr11 { 
		for idx2, val2 := range val {
			fmt.Println(idx, idx2, "<索引与值>", val2)
		}
	}
}
