package main

import (
	"fmt"
	"reflect"
)

// 练习

func main() {
	var aa = 3 / 2 // !=1.5 但却 =1, int型

	fmt.Println("3/2 =", aa, reflect.TypeOf(aa), "\n") // =1，int型

	var bb = float32(3) / 2 //float32型
	fmt.Println("float32(3) / 2 = ", bb, "\n")

	fmt.Println("3/1.5 =", 3/1.5, reflect.TypeOf(3/1.5)) // =2，float64型
	fmt.Println("3/1.0 =", 3/1.0, reflect.TypeOf(3/1.0)) // =3，float64型
	fmt.Println(
		reflect.TypeOf(aa), "\n",
		reflect.TypeOf(bb),
	)

	fmt.Println("两个整数得整数：", 3/2, 3*2)
	fmt.Println("两个浮点得浮点：", 1.5*2.0, reflect.TypeOf(1.5*2.0))
	fmt.Println("有一浮点得浮点：", 3.0/1, 3/1.0, 3.0*1, 3*1.0) //通过 reflect.TypeOf() 它们便是 float64

	var n01 = 3
	fmt.Println(
		n01*1.0/2,          //1 int
		1.0*n01/2,          //1 int
		0.5*2*n01,          //3 int
		0.6*2*float64(n01), // 3.5999999999999996（小数无法与整数相乘，报错！）
	)
	/*
		几乎所有字面量（literal）都是untyped 【无类型变量】
		如有一个【无类型变量】跟一个【有类型变量】运算时
		会转换到【有类型变量】的类型
	*/

	fmt.Println(
		3*1.0/2,
		3.0*1/2,
		3*1/2.0,
	)
	// 结果都是1.5，因为如果运算对象都是 untyped，运算结果按一定优先顺序
}
