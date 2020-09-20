// +build ignore

package main
import (
	"fmt"
	"strings"
)

func fn01 (sentence string) map[string]int {
	sentence = strings.ToLower(sentence)
	result := map[string]int{}
	strSlice := strings.Split(sentence, " ")
	fmt.Println(strSlice)
	// do something
	for _, curVal := range strSlice {
		if curVal == "" {
			continue
		}
		_, exist := result[curVal]
		if exist {
			result[curVal]++
		}else{
			result[curVal] = 1
		}
	}
	return result
}

func main() {
	var sentence01 = "How do    you do, how are you?"
	var result01 = fn01(sentence01)
	fmt.Println(result01)
}