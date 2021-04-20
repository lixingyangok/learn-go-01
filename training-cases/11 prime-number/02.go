// +build ignore

// 2021.04.20 20:49:1 星期二

package main

import "fmt"

func checkFn(num int) bool {
	if num <= 1 {
		return false
	}
	for idx := 2; idx < num; idx++ {
		if num % idx == 0 {
			return false
		}
	}
	return true
}

func toPrintThem(start, end int) {
	total := 0
	fmt.Printf("结果：\n")
	for idx := start; idx < end; idx++ {
		if checkFn(idx) {
			total++
			fmt.Printf("%d、", idx)
		}
	}
	fmt.Printf("\n质数：%d 个\n", total)
}

func main() {
	toPrintThem(0, 100)
}

