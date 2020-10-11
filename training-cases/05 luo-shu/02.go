// +build ignore

package main

import (
	"fmt"
)

func main() {
	var slice01 = [3][3]int{}
	var nineNumber = [...]int{4, 9, 2, 3, 5, 7, 8, 1, 6}
	for idx, cur := range nineNumber {
		size := 3                // 每行3个成员
		order := idx + 1         // 自然序号
		arrIndex := func() int { //小组号
			result := order / size
			if (result > 0) && (order%size == 0) {
				return result - 1
			}
			return result
		}()
		positionIdx := func() int { //位置号
			result := order % size
			if result == 0 { // =0 表示在末位
				return size - 1
			}
			return result - 1 // -1 表自然序转索引
		}()
		slice01[arrIndex][positionIdx] = cur
	}
	fmt.Println(slice01[0])
	fmt.Println(slice01[1])
	fmt.Println(slice01[2])
}
