// +build ignore

// 2021.05.05 12:31:46 星期三

package main

import "fmt"

func doer(peopleNumber int) int {
	times := 0
	for idx := peopleNumber - 1; idx > 1; idx-- {
		times += idx
	}
	return times
}

// 计算握手次数
func main() {
	times := doer(5)
	fmt.Printf("次数：%d", times)
}
