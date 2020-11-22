// +build ignore

package main

import (
	"fmt"
)

func main() {
	slice01 := []int{5, 9, 1, 7, 3}
	fmt.Println("原始值：", slice01)
	for i01 := 0; i01 < len(slice01); i01++ {
		for i02 := i01 + 1; i02 < len(slice01); i02++ {
			if slice01[i01] > slice01[i02] {
				slice01[i01], slice01[i02] = slice01[i02], slice01[i01]
			}
		}
		fmt.Println("slice01", slice01)
	}
	fmt.Println("结果值：", slice01)

	// 小黄  数学成绩100 满分120
	// 需要跟现在班级10个同学做比对
	// 他们的分数分别是  101 102 103 104 105 95 96 97 98 99
	// 按照  跟小明的“分数差”做对比 进行排序 如果分数差相等  则分数较小的排在前面
	// 差值越大越在前
}
