// +build ignore

package main

import (
	"fmt"
)

func main() {
	slice01 := []int{1, 7, 3, 5, 2, 9, 6}
	for idx, cur := range slice01 {
		if idx+1 == len(slice01) {
			break
		}
		next := slice01[idx+1]
		// curVal := cur
		if cur > next {
			slice01[idx] = next
			slice01[idx+1] = cur
		}
	}
	fmt.Println("slice01", slice01)
	// 小黄  数学成绩100    满分120
	// 需要跟现在班级10个同学做比对
	// 他们的分数分别是  101 102 103 104 105 95 96 97 98 99
	// 按照  跟小明的分数差做对比 进行排序 如果分数差相等  则分数较小的排在前面
	// 差值越大越在前
	// huang := 100
	// classMates := []int{101, 102, 103, 104, 105, 95, 96, 97, 98, 99}

	// for	idx, cur := range classMates {
	// 	gap := 0
	// 	if cur > huang {
	// 		gap = cur - huang
	// 	}else if cur < huang{
	// 		gap = huang - cur
	// 	}else{
	// 		continue
	// 	}
	// }
}
