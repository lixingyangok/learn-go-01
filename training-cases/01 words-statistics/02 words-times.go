package main

import (
	"fmt"
	"strings"
)

// 查看一句英文中有多少单词
const sentence01 = " hello world 你好 hello you   "
const sentence02 = "hi how are you, are you ok"

func fn02 (str string) (result map[string]int) {
	step01 := strings.Split(str, " ") //分隔
	result = make(map[string]int, len(step01))
	for _, val := range step01 {
		if len(val) == 0 { continue }
		var iAmount, ok = result[val]
		if ok {
			result[val] = iAmount + 1;
		}else {
			result[val] = 1;
		}
	}
	return
}

func main() {
	var test01 = fn02(sentence01)
	fmt.Println(test01)
}