// +build ignore

// 2021.04.19 20:56:53 星期一

package main

import (
	"fmt"
)

func toPrint() {
	for idx := 1; idx <= 9; idx++ {
		for i2 := 1; i2 <= idx; i2++ {
			fmt.Printf("%d*%d=%2d  ", i2, idx, i2 * idx)
		}
		fmt.Println()
	}
}

func main() {
	toPrint()
}
