package main

import (
	"fmt"
	"strings"
)

// 统计字符串中每个单词出现的次数

func main() {
	var sentence = "How do you do "
	var aWordsSlice = strings.Split(sentence, " ")
	var wordsMap = make(map[string]int, len(aWordsSlice))
	for _, cur := range aWordsSlice {
		var val, ok = wordsMap[cur]
		// TODO 设法判断是否为不可见字符
		// if !cur { 
		// 	continue;
		// }
		if ok {
			wordsMap[cur] = val + 1
		} else {
			wordsMap[cur] = 1
		}
	}
	// fmt.Println(sentence);
	// fmt.Println(aWordsSlice);
	fmt.Println(wordsMap);
}

