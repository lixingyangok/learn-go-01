// +build ignore

package main

// 练习goto语句
import (
	"fmt"
	// "strings"
)

func toPrint(max int) {
	start := 0
labelName01:
	if start <= max {
		fmt.Println(start)
		start++
		goto labelName01 // goto 可以去函数“内”的任意标签
	}
}

func main() {
	toPrint(5)
}
