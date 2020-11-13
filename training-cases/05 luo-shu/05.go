// +build ignore

package main

import (
	"fmt"
)

// 4, 9, 2, 3, 5, 7, 8, 1, 6

func main() {
	slice01 := []int{4, 9, 2, 3, 5, 7, 8, 1, 6}
	slice02 := [][]int{}
	for idx:=0; idx<=2; idx++ {
		slice02 = append(slice02, slice01[idx*3: idx*3+3])
	}
	for _, curSpan := range slice02 {
		fmt.Println(curSpan)
	}
}
