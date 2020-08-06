package main

import "fmt"
import "strings"

var rules = map[string]int{
	"赵": 1,
	"钱": 2,
}

func fn01(total int, names []string) (
	map[string]int, int,
) {
	var result = map[string]int{}
	for _, curName := range names {
		var num = 0
		for key, val := range rules {
			if strings.Contains(curName, key) {
				num = val
				break
			}
		}
		result[curName] = num
		total -= num
	}
	return result, total
}

func main() {
	var result, rest = fn01(9, []string{"赵一", "钱二", "赵开心"})
	var result02, rest02 = fn01(5, []string{"孙三", "李四"})
	fmt.Println(result, rest)
	fmt.Println(result02, rest02)
}
