/*
 * @Author: 李星阳
 * @Date: 2020-04-22 13:36:09
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-05-02 13:32:40
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
	var int01 = 0; //默认为int型，按设备，可能是int32，int64
	name = "Tom"
	msg := "" //短声明只能用在函数中
	msg = "how are you!"
	fmt.Printf("Hi %s, %s, %s\n", name, msg, int01)
	fmt.Print("I'm fine thank you, and you?\n")
}
