//+build wireinject"：
package main

import (
	"fmt"
	"strings"
)

// 统计字符串中每个单词出现的次数
// 函数要明示其返回值类型
func fn01(str string) map[string]int {
	var aWordsSlice = strings.Split(str, " ")
	var wordsMap = make(map[string]int, len(aWordsSlice))
	for _, cur := range aWordsSlice {
		var val, ok = wordsMap[cur]
		if ok {
			wordsMap[cur] = val + 1
		} else {
			wordsMap[cur] = 1
		}
	}
	return wordsMap;
}

func main() {
	var result01 = fn01("How do you do")
	var result02 = fn01("Fine thank you, and you")
	fmt.Println(result01)
	fmt.Println(result02)
}
