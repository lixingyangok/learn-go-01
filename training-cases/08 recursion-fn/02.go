// +build ignore

package main

import (
	"fmt"
)

// ▼递归求和
func getSum(num int) int {
	if num > 0 {
		return num + getSum(num - 1)
	}else{
		return 0
	}
}

func main() {
	result := getSum(3)
	fmt.Println(result)
}
