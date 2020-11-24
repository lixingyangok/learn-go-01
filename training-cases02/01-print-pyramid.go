// +build ignore

package main

import (
	"fmt"
	"strings"
)

// 打印一个金字塔

func toPrint(symbol string, num int) {
	for idx := 1; idx <= num; idx++ {
		oneLine := strings.Repeat(symbol, idx)
		fmt.Println(oneLine)
	}
}

func main() {
	toPrint("★", 3)
	fmt.Println("--------------")
	toPrint("◆", 4)
}
