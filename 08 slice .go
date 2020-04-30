package main

import "fmt"

func main() {
	var a01 = [...]int{0, 1, 2, 4, 5}
	var s01 = a01[1:3]
	fmt.Println(s01, len(s01), cap(s01))

	var s02 = []string{"北京", "上海", "广州"}
	s02= append(s02, "深圳") //切片添加一个成员
	fmt.Println(s02)

}
