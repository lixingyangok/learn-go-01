package main

import "fmt"

func main() {
	fmt.Println("\n\n\n")
	var a01 [3]int //声明时要指定：长度、成员类型 成员默认为零值：false, 0, "",
	var a02 [3]bool
	var a03 = [3]bool{true, true}        //第3个成员是false
	var a04 = []int{0, 1, 2, 3, 4, 5, 6} //自动长度
	var a05 = [3]int{0: 10, 2: 30}       //自动长度

	fmt.Println(a01) // [0 0 0]
	fmt.Println(a02) // [false false false]
	fmt.Println(a03) // [true true false]
	fmt.Println(a04) // [0 1 2 3 4 5 6]
	fmt.Println(a05) // [10 0 30]
	fmt.Println("------------------------\n")

	var a11 [2][2]int //一个长度为2的数组，成员为长度为3的数组 //[[0 0 0] [0 0 0]]
	var a12 = [][2]string{ //长度可省略
		[...]string{"A", "B"}, //长度不能省略，至少要写入...
		[...]string{"C", "D"},
	}
	fmt.Println(a11)
	fmt.Println(a12)
	fmt.Println("遍历这个数组")
	for idx:=0; idx<len(a12); idx++ {
		fmt.Println(a12[idx])
	}
	for idx, val := range a12 {
		for idx2, val2 := range val {
			fmt.Println(idx, "-", idx2, "-", val2)
		}
	}
}

