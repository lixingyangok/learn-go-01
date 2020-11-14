// +build ignore

// 2020.10.19 10:34:3 星期一

package main

import (
	"fmt"
)

func main() {
	for idx := 1; idx <= 9; idx++ {
		for i2 := 1; i2 <= idx; i2++ {
			// ▼ %02d 表示不足2位前方补0
			// fmt.Printf("%d*%d=%02d  ", i2, idx, i2*idx)
			// ▼ 不足2位时补空格
			fmt.Printf("%d*%d=%2d  ", i2, idx, i2*idx)
		}
		fmt.Println()
	}
}
