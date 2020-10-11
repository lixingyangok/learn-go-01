// +build ignore

package main

import (
	"fmt"
)

func main() {
	var slice01 = [3][3]int{}
	var nineNumber = [...]int{4, 9, 2, 3, 5, 7, 8, 1, 6}
	var size = 3
	for idx := 0; idx < size; idx++ {
		thisPart := nineNumber[idx*size : idx*size+3]
		// slice01[idx] = thisPart
		// ▲赋值失败：cannot use nineNumber[idx * size:idx * size + 3] (type []int) as type [3]int in assignment
		// ▼赋值成功
		copy(slice01[idx][:], thisPart)
		// fmt.Println(idx, thisPart)
	}
	fmt.Println(slice01[0])
	fmt.Println(slice01[1])
	fmt.Println(slice01[2])
}
