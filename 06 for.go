/*
 * @Author: 李星阳
 * @Date: 2020-04-23 13:20:43
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-05-08 13:53:04
 * @Description: 
 */
package main

import "fmt"

const sLine = "-----------------------------------------------"
const sLine02 = "\n" + "-----------------------------------------------"
const s01 = "你好阳光"

var arr01 = [4]string{"春", "夏", "秋", "冬"}

func main() {
	fmt.Println("\n\n\n")
	for idx := 0; idx < 3; idx++ {
		fmt.Print(idx, " ")
	}
	fmt.Println(sLine02)

	for idx, cur := range s01 {
		fmt.Print(idx, " - ", cur, " - ") //中文的索引是不准的
		fmt.Printf("%c\n", cur)           //能正常打印 byte 和 rune 类型
	}
	fmt.Println(sLine)

	for idx := 0; idx < len(arr01); idx++ { //为什么这个会正常打印？
		fmt.Print(arr01[idx])
	}
}
