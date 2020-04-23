package main

import "fmt"

func main() {
	fmt.Println("\n\n\n")
	var a01 [3]int //声明时要指定：长度、成员类型 成员默认为零值：false, 0, "",
	var a02 [3]bool
	var a03 = [3]bool{true, true}        //第3个成员是false
	var a04 = []int{0, 1, 2, 3, 4, 5, 6} //自动长度
	var a05 = [3]int{0: 10, 2: 30}       //自动长度

	fmt.Println(a01)
	fmt.Println(a02)
	fmt.Println(a03)
	fmt.Println(a04)
	fmt.Println(a05)
	fmt.Println("------------------------\n\n")

	var a11 [2][3]int //一个长度为2的数组，成员为长度为3的数组
	fmt.Println(a11)
}
