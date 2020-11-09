// +build ignore

package main

import (
	"fmt"
)

// 接收不定参
// 注：...在前表示合并为切片，...在后表示解散
func getSum(number ...int) int {
	// 如果只收到一个参数，仍然会将参组成切片
	fmt.Println(number)
	result := 0
	for _, val := range number {
		result += val
	}
	return result
}

func main() {
	// 函数返回值可匿名，可具名
	// 具名返回值可简写 return
	// 可变参函数，即函数数量不固定
	// 固定参与可变参结合的函数
	// 函数返回一个函数
	// 闭包
	value := getSum(1, 2, 3)
	fmt.Println(value)

}
