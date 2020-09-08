// +build ignore

/*
 * @Author: 李星阳
 * @Date: 2020-08-09 14:06:09
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-09-08 19:58:54
 * @Description:
 */
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func stringToSlice(str string) (result []string) {
	reg01 := regexp.MustCompile(`[a-z]+`)
	resultByRegexp := reg01.FindAllStringSubmatch(str, -1) //取得与正则匹配的文本（切片
	result = []string{}                                    //声明一个切片（结尾处有{}则表示同时声明内存
	for _, cur := range resultByRegexp {
		result = append(result, cur[0])
	}
	return
}

func fn05(str string) (result map[string]int) {
	str = strings.ToLower(str) //转小写
	wordSlice := stringToSlice(str)
	result = make(map[string]int, len(wordSlice))
	for idx, cur := range wordSlice {
		fmt.Println(idx, cur)
		amount, ok := result[cur]
		val := 1
		if ok {
			val = amount + 1
		}
		result[cur] = val
	}
	return
}

func main() {
	const sentence01 = "	 hellO world 你好 hEllo You   "
	const sentence02 = "hi how are yoU, are yOu ok"
	result01 := fn05(sentence01)
	result02 := fn05(sentence02)
	//
	fmt.Println(result01)
	fmt.Println(result02)

}
