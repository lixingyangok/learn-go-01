// +build ignore

package main

import (
	"fmt"
)

type MyStruct struct {
	Name   string
	Age    int
	number int
}

func main() {
	// 无法打印
	// fmt.Println("MyStruct", MyStruct) // type Struct01 is not an expression
	struct01 := MyStruct {
		Name: "张三",
		Age: 15,
		number: -2,
	}
	fmt.Printf("struct01：%v\n", struct01)
	fmt.Printf("struct01：%V\n", struct01)
	fmt.Printf("struct01：%+V\n", struct01)
}
