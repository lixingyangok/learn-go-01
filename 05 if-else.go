package main

import "fmt"

func main() {
	// var age = 17
	// ▼将age声明到这个判断作用域内
	if age := 10+7; age >= 18 {
		fmt.Println(age, "开始赌博------")
	} else if age > 10 {
		fmt.Println(age, "即将赌博---")
	} else {
		fmt.Println(age, "不能赌博-")
	}
}