package main

import (
	"fmt"
	"strings"
)

func fn01(str string) (result map[string]int) {
	strSlice := strings.Split(str, " ")
	result = make(map[string]int, len(strSlice))
	
	for _, val := range strSlice {
		if len(val) == 0 {
			continue
		}
		times, ok := result[val]
		if ok {
			result[val] = times + 1
		} else {
			result[val] = 1
		}
	}
	return result
}

func main() {
	// 查看一句英文中有多少单词
	const sentence01 = "	 hello world 你好 hello you   "
	const sentence02 = "hi how are you, are you ok"

	result01 := fn01(sentence01)
	result02 := fn01(sentence02)
	fmt.Println(result01)
	fmt.Println(result02)
}
