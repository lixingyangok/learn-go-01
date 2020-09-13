package main
// 123+build wireinject"：


import "fmt"

// 函数外只能放标识符，即：声明变量、声明常量、声明函数、声明类型，等各种声明
var helloWorld = "hello world" // 合法声明

/*
fmt.Println()	//此为非法语句，不能放在这里（主函数外）
num01 := 100	//短声明只能放在主函数内
*/

func main() {
	fmt.Println(helloWorld)
	
}
