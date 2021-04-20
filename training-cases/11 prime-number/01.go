// +build ignore

// 2021.04.20 20:48:54 星期二
package main

import "fmt"

func checkIsItPrimeNumber(num int) bool {
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

func main() {
	start := 0
	end := start + 100
	for idx := start; idx <= end; idx++ {
		if checkIsItPrimeNumber(idx) {
			fmt.Printf("%d、", idx)
		}
	}
}
