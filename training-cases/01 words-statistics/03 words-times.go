// +build ignore

/*
 * @Author: 李星阳
 * @Date: 2020-08-07 20:40:41
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-09-08 19:53:58
 * @Description:
 */
package main

import (
	"fmt"
	"strings"
)

func fn03(str string) (result map[string]int) {
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

	result01 := fn03(sentence01)
	result02 := fn03(sentence02)
	fmt.Println(result01)
	fmt.Println(result02)
}
