// +build ignore

package main

import "fmt"

// func find(arr []interface{}) interface{} {
// 	return arr[1]
// }

func main() {
	a01 := [3]int{}
	a02 := [...]int{1,2,3}
	fmt.Printf("数组1 %+v\n", a01)
	fmt.Printf("数组2 %+v\n", a02)
}
