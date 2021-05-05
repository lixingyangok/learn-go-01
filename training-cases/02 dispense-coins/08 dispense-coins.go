// +build ignore

// 2021.05.05 12:10:36 星期三
package main

import (
	"fmt"
	// "strings"
)

var schedule = map[string]int{
	"张": 3,
	"李": 4,
}

func doer(people []string, total int)(
	map[string]int,
	int,
){
	table := map[string]int{}
	iRest := total
	for _, curPerson := range people {
		surname := string([]rune(curPerson)[0])
		number, exist := schedule[surname]
		if exist {
			table[curPerson] = number
			iRest -= number
		}else{
			table[curPerson] = 0
		}
	}
	return table, iRest
}

func main() {
	table, iRest := doer([]string{
		"赵一",
		"钱二",
		"孙三",
		"张家辉",
		"李白",
	}, 20)
	fmt.Printf("方案：%v\n", table)
	fmt.Printf("剩余：%d\n", iRest)
}
