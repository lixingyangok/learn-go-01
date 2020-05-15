package main

import "fmt"

// func find(arr []interface{}) interface{} {
// 	return arr[1]
// }

func main() {
	// 1维数组
	// 无值初始化：
	// 有值初始化、间隔赋值，自动长度
	var a01 = [3]int{}
	var a02 = [3]int{1, 2, 3}
	var a03 = [3]int{0: 100, 2: 200}
	var a04 = []int{1, 2, 3}
	// var a05 = []string{"辽宁", "山东"}
	fmt.Println(
		"\n", a01,
		"\n", a02,
		"\n", a03,
		"\n", a04,
		"\n", //find(a05),
	)

	// 下标取值，倒数取值，取一段
	fmt.Println(
		"\n", a02[1],
		"\n", a02[len(a02)-1], //最后一个
		"\n",
	)

	// 2维数组
	// 数组1：2位长度的数组，成员也是2位长度数组且内容为整形
	var a11 = [2][2]int{}
	var a12 = [2][2]int{
		{1, 2}, {3, 4},
	}
	fmt.Println(
		a11,
		a12,
	)

	// 数组成员求和
	var result = 0
	for _, val := range a02 {
		result += val
	}
	fmt.Println("数组之合：", result)

	// ▼查看哪两个数组成员之合等于8
	var arr21 = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var arr22 = [2]int{}
	for idx, val := range arr21 {
		if val >= 8 { continue }
		for idx2 := idx; idx2 < len(arr21); idx2++{
			
		}
	}
}
