// +build ignore

package main

import (
	"fmt"
)

// 4, 9, 2, 3, 5, 7, 8, 1, 6

func main() {
	slice01:= []int{4, 9, 2, 3, 5, 7, 8, 1, 6}
	slice02:= [][]int{}
	for idx := 0; idx < 3; idx++ {
		start := idx*3
		thisPiece := slice01[start:start+3]
		slice02 = append(slice02, thisPiece);
	}
	for idx, cur := range slice02 {
		fmt.Println(idx, cur)
	}
}
