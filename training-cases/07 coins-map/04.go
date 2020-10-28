// +build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func aaaTestDemo2() {
	// t *testing.T
	i, _ := strconv.ParseUint("00000000", 2, 8)
	max, _ := strconv.ParseUint("11111111", 2, 8)
	fmt.Println("count:", max+1)
	for i <= max {
	   res := fmt.Sprintf("%08s", strconv.FormatUint(i, 2))
	   res = strings.Replace(res, "1", "●", -1)
	   res = strings.Replace(res, "0", "○", -1)
	   fmt.Println(res)
	   i++
	}
}
func main() {
	aaaTestDemo2()
}
