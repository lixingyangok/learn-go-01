// +build ignore

// 2021.05.10 19:41:55 星期一

package main

import "fmt"

func fn01(people int) (result int) {
	result = 0
	for times := people - 1; times > 0; times-- {
		result += times
	}
	return result
}

func main() {
	result := fn01(4) // 3、2、1
	fmt.Printf("次数：%d\n", result)
}
