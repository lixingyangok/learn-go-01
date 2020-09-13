// +build ignore

package main

import "fmt"

/*
	▼ var 批量声明变量
	● 在函数体外声明的变量称之为全局变量
	● 全局变量只需要在一个源文件中定义，就可以在所有源文件中使用
	● 当然，不包含这个全局变量的源文件，需要使用“import”关键字引入全局变量所在的源文件之后才能使用这个全局变量。
*/
var (
	name string
	age  int
)

func main() {
	/*
		函数内部变量，是局部变量
		局部变量不是一直存在的，它只在定义它的函数被调用后存在，函数调用结束后这个局部变量就会被销毁。
	*/
	var int01 = 0 //默认为int型，按设备，可能是int32，int64
	name = "Tom"
	msg := "" //短声明，只能用在函数中
	msg = "how are you!"
	fmt.Printf("Hi %s, %s, %s\n", name, msg, int01)
	fmt.Printf(aa, bb)
	fmt.Print("I'm fine thank you, and you?\n")

	var num01 = 100
	var f01 = .1 // 省略小数点前后数值
	var f02 = 1.
	fmt.Print(num01, f01, f02)
}
