// +build ignore

package main

import (
	"fmt"
)

func showList(letters []string, long int) (result []string) {
	slice := [][long]int{}
	for idx := 0; idx < long; idx++{
		slice[len(slice)] = []
		// for idx, cur := range slice {
		// 	result[idx]
		// }
	}
}


func main() {
	v1 := showList([]string{"a", "b", "c"}, 3)
	fmt.Println(v1)
}
