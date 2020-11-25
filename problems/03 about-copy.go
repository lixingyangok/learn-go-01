// +build ignore

package main

import "fmt"

// func getArr(num int)

func main() {
	a1 := []int{1, 2}
	b1 := []int{100, 200, 300}
	backVal := copy(a1, b1)
	fmt.Println("backVal:", backVal) // 返回2（表示有2个值被复制到了新切片中
	fmt.Println("a1:", a1)           // [100 200]
	// 结论：copy 长到短，目标不长长

	a2 := []int{100, 200, 300}
	b2 := []int{1, 2}
	copy(a2, b2)
	fmt.Println("a2", a2) // [1 2 300]
	// 结论：copy 短到长，目标不缩短

	a3 := []int{100, 200, 300}
	b3 := []int{1, 2}
	a3 = b3[:]            // a3变短，彻底等于b3
	fmt.Println("a3", a3) // [1 2]
	a3[0] = 123
	fmt.Println("a3，b3", a3, b3) // [123 2] [123 2]
	// 结论：[:]复制后的值，其实是引用类型，因为复制后一改俱改

}
