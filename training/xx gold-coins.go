// +build ignore

package main

import "fmt"

var coins = 100
var nameMap = map[string]int{
	"Monday":    0,
	"Tuesday":   0,
	"Wednesday": 0,
	"Thursday":  0,
	"Friday":    0,
	"Saturday":  0,
	"Sunday":    0,
}
var rule = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
}
func getHow(key string) (how int) {
	for _, curLetter := range key {
		fmt.Println(string(curLetter))
	}
	return;
}

func main() {
	// fmt.Println(coins)
	// fmt.Println(nameMap)
	// fmt.Println(rule)
	for key, _ := range nameMap {
		getHow(key)
		// fmt.Printf("%T  ", key)
		// iHowMany, ok := rule[key]
		// if ok != nil {
		// 	continue
		// }
		// cur[key] = cur[key] + 1
		// coins--;
		// fmt.Println(key, cur)
	}
	// fmt.Println(nameMap)
	
}
