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

func (this *MyStruct) SayMyName() {
	fmt.Println("SayMyName：", this.Name)
}

func main() {
	// 无法打印
	// fmt.Println("MyStruct", MyStruct) // type Struct01 is not an expression
	struct01 := MyStruct{
		Name:   "张三",
		Age:    15,
		number: -2,
	}
	// ▼注意3种打印方法
	fmt.Printf("struct01：%v\n", struct01)
	fmt.Printf("struct01：%V\n", struct01)
	fmt.Printf("struct01：%+V\n", struct01)
	// 
	struct01.SayMyName()
	fmt.Printf("打印Name：%s\n", struct01.Name) //顺利打印
	fmt.Printf("打印number：%d\n", struct01.number) //顺利打印

}
