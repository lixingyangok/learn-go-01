/*
 * @Author: 李星阳
 * @Date: 2020-04-22 12:09:49
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-04-30 12:44:56
 * @Description: 
 */

package main

import "fmt"

// 函数外只能放标识符，即：声明变量、声明常量、声明函数、声明类型，等各种声明
// fmt.Println  // 此为非法语句，不能放在这里（主函数外）
var helloWorld = "hello world" // 合法声明

func main() {
	fmt.Println(helloWorld)
}
