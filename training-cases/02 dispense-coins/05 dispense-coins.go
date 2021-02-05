// +build ignore

// 2021.02.05 11:29:8 星期五
package main

import (
	"fmt"
	// "strings"
)

var schemeMap = map[string]int{
	"张": 1,
	"王": 2,
	"李": 3,
	"赵": 4,
}

func toAllot(total int, people []string) (
	rest int, resultMap map[string]int,
) {
	rest = total
	resultMap = map[string]int{}
	for _, curName := range people {
		firstLetter := string([]rune(curName)[0])
		_, exist := schemeMap[firstLetter]
		if exist {
			resultMap[curName] = schemeMap[firstLetter]
			rest -= schemeMap[firstLetter]
		} else {
			resultMap[curName] = 0
		}
	}
	return
}

func main() {
	rest, resultMap := toAllot(20, []string{
		"张三",
		"李四",
		"张仨",
		"李大",
		"刘大",
		"刘二",
	})
	fmt.Printf("剩余：%d \n", rest)
	fmt.Printf("方案：\n%+v\n", resultMap)
}
