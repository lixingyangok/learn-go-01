package main

import "fmt"

func main() {
	var a01 = [...]int{0, 1, 2, 3, 4} //数组
	var s01 = a01[2:]                 //从数组生成切片。[]内:符号，左包含右不包含
	var s02 = []int{}                 //新建空切片
	var s03 []int                     //没申请内存
	fmt.Println(
		s01,
		len(s01), //切片长度
		cap(s01), //切片容量
		"\n",
	)
	fmt.Println(
		s02, "\n",
		s03, "\n",
	)

	fmt.Println(
		a01[:], "\n", //从头取到尾
		a01[2:], "\n", //2参不传取到最后
		a01[:2], "\n", //1参不传从头取
	)

	var s02 = []string{"北京", "上海", "广州"}
	s02 = append(s02, "深圳") //切片添加一个成员
	fmt.Println(s02)

}
