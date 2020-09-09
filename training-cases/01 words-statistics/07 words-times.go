// +build ignore

// 2020.09.09 20:10:53 星期三
package main

import (
	"fmt"
	"strings"
)

func fn01(str string) (resultMap map[string]int) {
	lowerStr := strings.ToLower(str)
	strArr := strings.Split(lowerStr, " ")
	resultMap = make(map[string]int, len(strArr))
	for _, val := range strArr {
		_, exist := resultMap[val]
		if exist {
			resultMap[val]++
		} else {
			resultMap[val] = 1
		}
	}
	return
}

func main() {
	fmt.Println(
		fn01("Hello world hello you"),
	)
}
