package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 声明空map
	// 声明空map，指定长度，用make
	// 声明同时初始化map
	// 判断某键是否存在
	var m00 map[int]string
	var m01 = map[int]string{}
	var m02 = make(map[int]string, 8) //有8个键值对？
	fmt.Println(m00, m00 == nil)      // 是nil
	fmt.Println(m01, m01 == nil)      // 不是nil
	fmt.Println(m02, m02 == nil)      // 不是nil

	// m00[1] = "wrong" //报错，因为没有初始化
	m01[1] = "OK"
	m02[1] = "OK"
	fmt.Println(m01, m01 == nil) // 非nil
	fmt.Println(m02, m02 == nil) // 非nil

	// 声明同时初始化
	var m11 = map[int]string{
		1: "Tom",
		2: "Jack",
	}
	m11[3] = "Lucy"
	fmt.Println(m11)

	// 查看某键是否在map中
	val, ok := m11[3]
	val2, ok2 := m11[4]
	fmt.Println(ok, val)
	fmt.Println(ok2, val2)

	// 遍历一个map
	for key, val := range m11 {
		fmt.Println(key, val)
	}
	delete(m11, 1)
	fmt.Println(m11)

	// 通过遍历生成map
	var m12 = make(map[string]int, 5)
	for idx := 0; idx < 5; idx++ {
		key := fmt.Sprintf("stu%02d", idx)
		fmt.Println(key, rand.Intn(100))
		m12[key] = rand.Intn(100)
	}
	fmt.Println(m12)


	// 声明一个值为切片的map，然后不能立即赋值，如何解决？
	// 例：How do you do? 统计各个单词出现的数量
	
}
