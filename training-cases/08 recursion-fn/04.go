// +build ignore

package main

import (
	"fmt"
)

// 递归得到句子
func getSentence(input string) string{
	sentence := "从前有座山，山里有座庙，庙里有老道，大老道对小老道说■■■■■"
	if len(input) < 100 {
		return input + getSentence(input + sentence)
	}
	return input
}

func main() {
	v1 := getSentence("")
	fmt.Println(v1)
}
