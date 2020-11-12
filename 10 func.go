// +build ignore

package main

import (
	"fmt"
)

// 接收不定参
// 注：...在前表示合并为切片，
func getSum(number ...int) int {
	// 如果只收到一个参数，仍然会将参组成切片
	fmt.Println(number)
	result := 0
	for _, val := range number {
		result += val
	}
	return result
}

func closedPkg(step int) func() {
	var num = 0 // 闭包，这个变量是他们的账本
	return func() {
		num += step
		fmt.Println(num)
	}
}

func main() {
	value := getSum(1, 2, 3)
	fmt.Println(value) // 6

	// ▼复习 ...在后的情况
	slice01 := []int{1, 2, 3}
	slice01 = append(slice01, []int{4, 5, 6}...) // ...在后表示解散
	fmt.Println(slice01)                         // 1，2，3，4，5，6

	fn01 := closedPkg(10)
	fn02 := closedPkg(100)
	fn01()
	fn01()
	fn02()
	fn02()
}

// 函数返回值可匿名，可具名
// 具名返回值可把返回简写为 return
// 可变参函数，即函数数量不固定（组成切片）
// 固定参与可变参结合的函数
// 函数返回一个函数
// 闭包
