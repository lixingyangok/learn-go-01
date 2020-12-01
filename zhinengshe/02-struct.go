// +build ignore

package main

import "fmt"

func main() {
	type user struct {
		name    string
		age     uint8
		amount  float64
		friends []user
	}
	// ▲声明一个结构体

	// ▼按结构体格声明一个元素
	var user01 user
	user01.name = "Tom"
	user01.age = 18
	user01.amount = 1.23456789
	fmt.Printf("%+v\n", user01)

	// ▼声明方式2（常用方式？）
	var user02 = user{ //此处有类型推导，不必显示声明
		name: "Jack",
	}
	fmt.Printf("%+v\n", user02)
}
