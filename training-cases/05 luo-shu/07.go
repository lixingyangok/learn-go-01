// +build ignore

package main

import (
	"fmt"
)

// 4, 9, 2, 3, 5, 7, 8, 1, 6
var arr01 = [...]int{4, 9, 2, 3, 5, 7, 8, 1, 6}

func main() {
	fmt.Printf("数组：\n%+v\n", arr01)
	const oneLine = 3
	// slice01 := [][]int{} 
	slice01 := make([][]int, 3)
	for idx, curVal := range arr01 {
		lineNo := int(idx / oneLine)
		seatNo := idx % oneLine
		if len(slice01[lineNo]) == 0 { // 
			slice01[lineNo] = make([]int, 3)
		}
		slice01[lineNo][seatNo] = curVal
	}
	fmt.Printf("二维数组：\n%+v\n", slice01)
}
