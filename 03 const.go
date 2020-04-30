/*
 * @Author: 李星阳
 * @Date: 2020-04-22 13:36:09
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-04-28 13:52:21
 * @Description: 
 */
package main

import "fmt"

const (
	var01 = 100
	var02 //不声明值就和上一行一样（仅限常量
	var03 //和上一行一样
)
const (
	_ = iota // 0
	l01 = "插队者" //
	l02 = iota // 2
	l03 // 3
)
const (
	d1, d2 = iota + 10, iota + 10 // 10, 10
	d3, d4 = iota + 10, iota + 10 // 11, 11
)


func main() {
	fmt.Println(var02, var03)
	fmt.Println(l01, l02, l03)
	fmt.Println(d1, d2, d3, d4)
}

