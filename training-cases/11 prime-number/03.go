// +build ignore

// 2021.04.20 20:49:1 星期二

package main

import "fmt"

func checker(num int) bool {
	if num < 2 {
		return false
	}
	for idx := 2; idx < num; idx++ {
		if num%idx == 0 {
			return false
		}
	}
	return true
}

func printer(n1, n2 int) {
	for n1 < n2 {
		if checker(n1) {
			fmt.Printf("%d、", n1)
		}
		n1++
	}
}

func main() {
	printer(1, 99)
}
