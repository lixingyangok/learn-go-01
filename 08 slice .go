// +build ignore

package main

import "fmt"

func main() {
	var a01 = [...]int{0, 1, 2, 3, 4} //数组
	var s01 = a01[2:]                 //从数组生成切片。[]内:符号，左包含右不包含
	var s02 = []int{}                 //【空切片】已申请内存
	var s03 []int                     //【空切片】但没申请内存，即nil，即 s03 == nil
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
	var s04 = []string{"北京", "上海", "广州"}
	var s05 = make([]string, len(s04), 99) // 通过 make 函数申请内存
	var s05b = []string{} //有内存空切片
	s04 = append(s04, "深圳") //切片添加一个成员
	copy(s05, s04) //长度4的切片，覆盖给长度3的切片，最终长度=3
	s05b = append(s05b, s04...)
	s05b[0] = "南京"
	fmt.Println(s04)
	fmt.Println(s05)
	fmt.Println(s05b)

	var s06 = make([]string, len(s05b))
	s06 = append(s05b[:1], s05b[2:]...)
	fmt.Print(s06)
}


