// +build ignore

package main

import (
	"fmt"
	"strings"
)

// 2020.11.14 17:36:4 星期六

var myMap = map[string]int{
	"张": 3,
	"李": 4,
	"赵": 3,
	"马": 4,
	"王": 5,
}

func allocationCoins(
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
		for curFamilyName, num := range myMap {
			if strings.HasPrefix(curName, curFamilyName) {
				needToGive = num
				continue
			}
		}
		if needToGive > rest {
			result[curName] = rest
			rest -= rest
		} else {
			result[curName] = needToGive
			rest -= needToGive
		}
	}
	return rest, result
}

func main() {
	rest, result := allocationCoins(6, []string{"张三", "李四", "张甲", "李丁"})
	fmt.Println(rest)
	fmt.Println(result)
}
