/*
* @Author: 李星阳
* @Date: 2020-04-22 12:09:49
* @LastEditors: 李星阳
* @LastEditTime: 2020-04-22 13:43:24
* @Description: ain
*/
package main

import "fmt"

// 函数外只能放标识符：变量、常量、函数、类型的声明
// fmt.Println  // 非法语句，语句不能放在这里（函数外）
var helloWorld = "hello world" // 合法声明

func main() {
	fmt.Println(helloWorld)
}
