// +build ignore

package main

import (
	"fmt"
)

// 解决这个叫做 Fizz-Buzz 的问题：
// 编写一个程序，打印从 1 到 100 的数字。
// 当是3个倍数数就打印 “Fizz” 代替数字，
// 当是5的倍数就打印 “Buzz” 。
// 当数字同时是3和5的倍数 时，打印 “FizzBuzz” 。

func main() {
	for idx := 1; idx <= 100; idx++ {
		if idx%10==0{
			fmt.Println()
		}
		if idx%3 == 0 && idx%5 == 0 {
			fmt.Print("FizzBuzz ", idx, "、 ")
			continue
		}
		if idx%3 == 0 {
			fmt.Print("Fizz ", idx, "、 ")
		} else if idx%5 == 0 {
			fmt.Print("Buzz ", idx, "、 ")
		}
	}
}
