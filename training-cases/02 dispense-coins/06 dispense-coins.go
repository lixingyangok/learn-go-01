// +build ignore

// 2021.03.28 20:28:21 星期日
package main

import (
	"fmt"
	"strings"
)

var schemeMap = map[string]int{
	"张": 1,
	"王": 2,
	"李": 3,
	"赵": 4,
}

func toAllocate(coinsNumber int, people []string) (
	rest int,
){
	needToGive := 0
	planMap := map[string]int{}
	for _, curName := range people {
		for curFamilyName, giveNumber := range schemeMap {
			if !strings.HasPrefix(curName, curFamilyName) {
				continue
			}
			// fmt.Printf( "%s 由 %s 开头，分 %d 个\n", curName, curFamilyName, giveNumber, )
			planMap[curName] = giveNumber
			needToGive += giveNumber;
		}
	}
	fmt.Printf("方案：%v\n", planMap)
	rest = coinsNumber - needToGive
	return rest;
}



func main() {
	result 	:= toAllocate(10, []string{
		"张三", 
		"李四",
		"王五",
	})
	fmt.Printf("剩余：%d", result)
}
