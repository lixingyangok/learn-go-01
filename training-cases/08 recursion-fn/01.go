// +build ignore

package main

import (
	"fmt"
)

func printNumberToSmall(number int){
	if number > 0 {
		fmt.Print(number, "、")
		printNumberToSmall(number - 1)
	}
}

func printNumberToBig(n1, n2 int){
	fmt.Print(n1, "、")
	if n1 < n2 {
		printNumberToBig(n1 + 1, n2)
	}
}

func main() {
	// 递归打印1-5
	printNumberToSmall(5);
	fmt.Println("\n--------------------")
	printNumberToBig(1, 5);

	
	// 递归求和1-10
	// 递归5的阶乘
}
