// +build ignore

package main

import (
	"fmt"
)

func main() {
	// 关于变量命名：重复其实是可以的，但你不能没一个新的
	a, b := 11, 22
	a, c := 33, 44
	fmt.Println("a,b,c", a, b, c)

	x, z := "aa", "bb"
	y, z := "cc", "dd"
	fmt.Println("x,y,z", x, y, z)

	newVal01, a, b, c, x, y, z := 1, 2, 3, 4, "5", "6", "7", "8"
	fmt.Println("newVal01, a, b, c, x, y, z ------", newVal01, a, b, c, x, y, z)
}
