// +build ignore

package main

import (
	"fmt"
	"strings"
)
// 2020.09.14 8:7:36 星期一

var myMap = map[string]int{
	"张": 3,
	"李": 4,
	"王": 5,
	"赵": 6,
	"马": 9,
}

func fn01 (total int, theSlice []string) (map[string]int, int) {
	result := map[string]int{}
	for _, curName := range theSlice {
		needToGive := 0
		for curFamilyName, curNum := range myMap {
			if strings.HasPrefix(curName, curFamilyName) {
				needToGive = curNum
				continue
			}
		}
		result[curName] = needToGive
		total -= needToGive
	}
	return result, total
}

func main() {
	result, rest := fn01(25, []string{
		"张三",
		"张仨",
		"李四",
		"马英九",
		"王五",
		"赵六",
		"司马七",
		"诸葛八",
	})
	fmt.Println(rest)
	fmt.Println(result)
}
