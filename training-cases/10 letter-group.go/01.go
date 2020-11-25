// +build ignore

package main

import (
	"fmt"
)

func showList(letters []string, long int) (result [][]int) {
	result = [][]int{}
	slice := make([]int, long)
	for idx := 0; idx < long; idx++ {
		for i2 := 0; i2 + 1 < len(letters); i2++ {
			sliceForPush := []int{0,0,0}
			copy(sliceForPush, slice)
			result = append(result, sliceForPush)
			slice[idx]++
		}
	}
	return
}


func main() {
	v1 := showList([]string{"a", "b", "c"}, 3)
	fmt.Println(v1)
}
