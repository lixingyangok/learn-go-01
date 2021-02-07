package main

import (
	"fmt"
	"time"
)

// ▼函数外只能放“标识符”，即只能声明，不可执行语句，如不可执行打印语句
// ▼可以声明常量、变量、函数、类型等
const helloWorld = "Hello world"
var myStr = fmt.Sprintf("%s and hello you", helloWorld)
var dataString = func() string { // 此函数会执行
	now := time.Now()
	fmt.Println("★★★★") // 此处会打印
	return now.Format("2006.01.02 03:04:05 pm Mon")
}()

// ▼错误示例
// num01 := 100	// 短声明只能放在函数内
// fmt.Println()	//此为非法语句，不能放在这里（不能放主函数外）

func main() {
	fmt.Println("用 Println 打印：", helloWorld)
	fmt.Println("组合后的内容：", myStr)
	fmt.Println("今天日期信息：", dataString)
}

/*
	■■■■■■■■■ ▼数字系列 ■■■■■■■■■■■■■■■■■■■■■
	%d 打印整型数，十进制输出(如果d前面有数字，表示控制输出宽度，默认使用空白填充，%05d 会在不满5位时填充0)
	%b 打印整型数，二进制输出
	%o 打印整型数，八进制输出，如果x前面带有#表示带有零的八进制
	%x 打印整型数，十六进制输出，如果x前面带有#表示带有0x的十六进制
	%c 打印整型数，字符输出（如果有）
	%f 打印浮点数，正常输出（如果f前面有x.y 表示控制宽度和输出小数点位数，要对齐再加-，例如 %-10.5f）
	%e,%E 打印浮点数，科学记数法输出
	■■■■■■■■■ ▼字符串 ■■■■■■■■■■■■■■■■■■■■■
	%s 打印字符串
	%q 打印字符串（带引号—）
	%x 打印字符串（使用base-16输出其编码
	%U 打印Unicode字符
	%#U 打印带字符的Unicode
	■■■■■■■■■ ▼其它 ■■■■■■■■■■■■■■■■■■■■■
	%v 打印结构体
	%+v 打印结构体（键值对形式，数值带正号）
	%T 打印对象的类型
	%t 打印布尔值
	%p 打印指针
*/
