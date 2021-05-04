// +build ignore

// 2021.05.04 18:27:49 星期二
package main

import (
	"fmt"
	// "strings"
)

var schedule = map[string]int{
	"赵": 1,
	"钱": 2,
	"孙": 3,
	"李": 4,
	"张": 3,
}

func doer(people []string, sum int) (
	map[string]int,
	int,
) {
	result := map[string]int{}
	iRest := sum
	for _, cur := range people {
		surname := string(
			[]rune(cur)[0],
		)
		val, exist := schedule[surname]
		if exist {
			result[cur] = val
			iRest -= val
		}else{
			result[cur] = 0
		}
	}
	return result, iRest
}

func main() {
	result, rest := doer([]string{"张三", "李四", "王五"}, 20)
	fmt.Printf("方案：%v\n剩余：%d\n", result, rest)
}
