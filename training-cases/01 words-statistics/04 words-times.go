package main

import (
	"fmt"
	"strings"
)

func fn04(str string) (result map[string]int) {
	str = strings.Trim(str, " 	")
	var wordSlice = strings.Split(str, " ")
	result = make(map[string]int, len(wordSlice))
	for _, curWord := range wordSlice {
		if len(curWord) == 0 {
			fmt.Println("长度0")
			continue
		}
		var amount, ok = result[curWord]
		if ok {
			result[curWord] = amount + 1
		}else {
			result[curWord] = 1
		}
	}
	return
}

func main() {
	// 查看一句英文中有多少单词
	const sentence01 = "	 hello world 你好 hello you   "
	const sentence02 = "hi how are you, are you ok"
	var result01 = fn04(sentence01);
	fmt.Println(result01)
}
