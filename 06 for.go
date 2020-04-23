package main

import "fmt"

func main() {
	s02 := "天天开心"
	for idx, cur := range s02 {
		fmt.Print(idx, " - ", cur,  " - ") //索引是不准的
		fmt.Printf("%c\n", cur) //能正常打印 byte 和 rune 类型
	}
}
