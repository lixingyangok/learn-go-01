// +build ignore

package main

import (
	"fmt"
	// "math"
)

func main() {
	/*
		var goArr01 = []int{} // 合法声明，但无法修改
		var goArr02 = make([]int, 10) // 合法声明，可以修改
		// goArr01[0] = 0
		goArr02[0] = 2
		fmt.Println(goArr01)
		fmt.Println(goArr02)
	*/
	// 4,9,2, 3,5,7, 8,1,6
	// var slice01 = [3][3]int{}
	var nineNumber = [...]int{4,9,2,3,5,7,8,1,6}
	for idx, _ := range nineNumber {
		size := 3
		order := idx + 1
		rest := order % size
		step01 := int(order / size) + 1
		if rest == 0 {
			step01--
			rest = size
		}
		step01--
		rest--
		fmt.Println(step01, rest)
		// slice01[step01][rest] = cur
	}
	// fmt.Println(slice01)
}
