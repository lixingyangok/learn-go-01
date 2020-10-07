// +build ignore

package main

import (
	"fmt"
	// "math"
)

func main() {
	/*
		var goArr01 = []int{} // 合法声明，但无法修改
		var goArr02 = make([]int, 10)
		// ▲合法声明，可以修改
		// ▲那就是说，要声明时要表明切片用于保存什么类型数据，表示多长
		goArr01[0] = 0 //报错
		goArr02[0] = 2 //可执行
		fmt.Println(goArr01)
		fmt.Println(goArr02)
	*/
	var slice01 = [3][3]int{}
	var nineNumber = [...]int{4, 9, 2, 3, 5, 7, 8, 1, 6}
	for idx, cur := range nineNumber {
		size := 3                       // 每行3个成员
		order := idx + 1                // 自然序号
		// placeIndex := order % size      // 座位号
		// if placeIndex == 0 {            // =0 表示在车辆的最后一个座
		// 	arrIndex--
		// 	placeIndex = size
		// }
		// placeIndex--
		arrIndex := func()int{
			result := int(order/size)
			if result % 1 == 0 {
				return int(result) - 1
			}
			return int(result)
		}()
		positionIdx := func() int { //位置号
			result := order % size
			if result == 0 {
				return size
			}
			return result - 1 // -1 表自然序转索引
		}()
		// arrIndex := int(order/size) + 1 // 数组序号
		// arrIndex--
		slice01[arrIndex][positionIdx] = cur
	}
	fmt.Println(slice01[0])
	fmt.Println(slice01[1])
	fmt.Println(slice01[2])

}
