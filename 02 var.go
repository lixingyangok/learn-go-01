/*
 * @Author: 李星阳
 * @Date: 2020-04-22 13:36:09
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-04-22 13:44:05
 * @Description: 
 */

package main

import "fmt"

// ▼批量声明变量
var (
	name string
	age  int
)

func main() {
	name = "Tom"
	msg := "" //短声明只能用在函数中
	msg = "how are you!"
	fmt.Printf("Hi %s, %s\n", name, msg)
	fmt.Print("I'm fine thank you, and you?\n")
}
