package main

import (
	"fmt"
	// "strconv"
	// "strings"
	"testing"
)

/* 
先 cd 到当前目标
然后执行：go test -v xx_test.go
可执行当前文件中的所有函数
*/

// ▼测试方法（不用显示的调用）当执行 go test 时会自动执行它
func TestDemo2(t *testing.T) {
	fmt.Println("123")
	// i, _ := strconv.ParseUint("00000000", 2, 8)
	// max, _ := strconv.ParseUint("11111111", 2, 8)
	// fmt.Println("count:", max+1)
	// for i <= max {
	// 	res := fmt.Sprintf("%08s", strconv.FormatUint(i, 2))
	// 	res = strings.Replace(res, "1", "●", -1)
	// 	res = strings.Replace(res, "0", "○", -1)
	// 	fmt.Println(res)
	// 	i++
	// }
}


func TestDemo1(t *testing.T) {
	fmt.Println("11111111")
	// i, _ := strconv.ParseUint("00000000", 2, 8)
	// max, _ := strconv.ParseUint("11111111", 2, 8)
	// fmt.Println("count:", max+1)
	// for i <= max {
	// 	res := fmt.Sprintf("%08s", strconv.FormatUint(i, 2))
	// 	res = strings.Replace(res, "1", "●", -1)
	// 	res = strings.Replace(res, "0", "○", -1)
	// 	fmt.Println(res)
	// 	i++
	// }
}