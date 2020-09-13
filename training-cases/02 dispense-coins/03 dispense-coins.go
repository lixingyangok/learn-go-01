// +build ignore

package main

import "fmt"

var plan = map[string]int{
	"张": 1,
	"李": 2,
}

func toAllocate(total int, nameSlice []string) (rest int, result map[string]int) {
	rest = total
	result = map[string]int{}
	for _, curName := range nameSlice {
		needToGive := 0
		for curFamilyName, numberToGive := range plan {
			if string([]rune(curName)[0]) == curFamilyName {
				needToGive = numberToGive
			}
		}
		result[curName] = needToGive
		rest -= needToGive
	}
	return
}

func main() {
	rest, result := toAllocate(10, []string{
		"张三",
		"张仨",
		"李四",
		"王五",
		"赵六",
	})
	fmt.Println(rest)
	fmt.Println(result)
}
