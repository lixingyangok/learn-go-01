// +build ignore

package main

// 2020.09.09 20:0:49 星期三

import (
	"fmt"
	"strings"
)

func fn01(str string) (resultMap map[string]int) {
	lowerString := strings.ToLower(str)
	strArr := strings.Split(lowerString, " ")
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
	var result01 = fn01("Hello world hello you")
	fmt.Println(result01)
}
