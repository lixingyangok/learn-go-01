// +build ignore

package main

import (
	"fmt"
)

func main() {
	// 注意：make的2参是0，如果非0，则生成的切片会有默认的成员
	// 3参可以省略
	var slice01 = make([][3]int, 0, 10) // 保存由0、1两个数字能形成的各种组合
	for idx := 0; idx <= 1; idx++ {
		for idx02 := 0; idx02 <= 1; idx02++ {
			for idx03 := 0; idx03 <= 1; idx03++ {
				// var newOne = [3]int{idx, idx02, idx03}
				// slice01 = append(slice01, newOne)
				// ▲注意上下2种插入方法▼
				slice01 = append(slice01, [3]int{idx, idx02, idx03})
			}
		}
	}

	var slice02 = make([]string, 0, 10) // 数组形式转字符形式
	var yang = "■■■■■■■"
	var yin = "■■■ ■■■"
	for _, curArr := range slice01 {
		var symbol = ""
		for idx, curVal := range curArr {
			var mark = yin
			if curVal == 1 {
				mark = yang
			}
			if idx+1 < len(curArr) {
				mark += "\n"
			}
			symbol += mark
		}
		slice02 = append(slice02, symbol)
	}
	// fmt.Println(slice01)
	// fmt.Println(slice02)
	for idx, curSymbol := range slice02 {
		fmt.Println("-", idx+1, "-------------------------------------")
		fmt.Println(curSymbol)
	}
}
