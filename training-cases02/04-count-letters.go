// +build ignore

package main

import (
	"fmt"
)

// 建立一个程序统计字符串里的字符数量： 
// 同时输出这个字符串的字节数。 提示： 看看 unicode/utf8 包。

const theString = "asSASA ddd dsjkdsjs dk 中文"

func fn01(theStr string) (result map[string]int){
	result = make(map[string]int, len(theStr) / 2)
	for _, cur := range []rune(theStr) {
		curLetter := string(cur)
		_, exist := result[curLetter]
		if exist {
			result[curLetter]++
		}else {
			result[curLetter] = 1
		}
	}
	return
}

func fn02(theStr, newStr string, start int) string {
	myRune := []rune(theStr)
	newRune := []rune(newStr)
	for idx, _ := range myRune {
		if idx >= start && idx <= start + len(newRune) - 1 {
			
			myRune[idx] = rune("工") // newRune[len(newRune) - idx]
		}
	}
	return string(myRune)
}

func main() {
	v1 := fn01(theString)
	v2 := fn02("abc这是一段文字xyz", "★●■", 2)
	fmt.Println(v1)
	fmt.Println(v2)
}
