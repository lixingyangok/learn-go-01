// +build ignore

package main

import (
	"fmt"
	"strings"
)
// 2020.11.14 17:36:4 星期六

var myMap = map[string]int{
	"赵": 3,
	"马": 4,
	"张": 6,
	"李": 9,
	"王": 5,
}

func allocationCoins (
	coinsTotal int, nameSlice []string, 
) (
	rest int, result map[string]int, 
) {
	// fmt.Println("nameSlice", nameSlice)
	// fmt.Println("coinsTotal", coinsTotal)
	result = map[string]int{}
	rest = coinsTotal
	for _, curName := range nameSlice {
		needToGive := 0
		for curFamilyName, num  := range myMap {
			if strings.HasPrefix(curName, curFamilyName) {
				needToGive = num
				continue
			}
		}
		result[curName] = needToGive
		rest -= needToGive
	}
	return rest, result
}

func main() {
	rest, result := allocationCoins(22, []string{"张三", "李四"})
	fmt.Println(rest)
	fmt.Println(result)
}
