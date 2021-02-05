// +build ignore

// 2021.02.05 14:22:37 星期五
package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string) (result map[string]int) {
	fmt.Printf("结果值：%p\n", result)
	result = map[string]int{} // 如果已经【声明了】返回值的变量名，此处要给此变量声明内存
	fmt.Printf("结果值：%p\n", result)
	// ■■ ▲注意这2次打印内存地址的不同 ■■■■■■■■■■■■■■■■■■■■■■■■■■■■
	// result := map[string]int{} // 如果【没声明】返回值的变量，更要在此做一个初始化
	slice01 := strings.Split(sentence, " ")
	for _, val := range slice01 {
		_, exist := result[val]
		if exist {
			result[val]++
		} else {
			result[val] = 1
		}
	}
	return result
}

func main() {
	result := countWords("hello world hello you")
	fmt.Printf("打印：%+v \n\n", result)
}
