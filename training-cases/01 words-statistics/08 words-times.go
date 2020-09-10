// +build ignore

// 2020.09.10 6:48:28 星期四
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		fn01("Hello world, hello you"),
	)
}

func fn01(str string) (resultMap map[string]int) {
	lowerStr := strings.ToLower(str)
	strSlice := strings.Split(lowerStr, " ")
	resultMap = make(map[string]int, len(strSlice))
	for _, val := range strSlice {
		_, exist := resultMap[val]
		if exist {
			resultMap[val]++
		} else {
			resultMap[val] = 1
		}
	}
	return
}
