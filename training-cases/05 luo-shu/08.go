// +build ignore

package main

import (
	"fmt"
)

// 4, 9, 2, 3, 5, 7, 8, 1, 6
var arr01 = [...]int{4, 9, 2, 3, 5, 7, 8, 1, 6}

func getSlice() [][]int {
	// ▼ 这样声明，然后赋值是行不通的
	// var result [][]int 
	// ▼ 这样可以
	var result = make([][]int, 3)
	for idx, _ := range result {
		// result[idx] = make([]int, 3) // 可执行
		result[idx] = arr01[idx*3 : idx*3+3]
	}
	return result
}

func main() {
	slice01 := getSlice()
	for idx, cur := range slice01 {
		fmt.Printf("%d %v\n", idx, cur)
	}
}

