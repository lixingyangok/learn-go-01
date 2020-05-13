package main

import (
	"fmt"
	"reflect"
)

func main() {
	var aa = 3 / 2 // =1, int型

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
	fmt.Println(n01 * 1.0 / 2) // 1     如果左侧运算符类型确定，则右侧转为左侧类型再运算
	fmt.Println(3 * 1.0 / 2)   // 1.5   如果左侧类型不确定，则根据右侧类型推导左侧类型
	/*
		几乎所有字面量（literal）都是untyped 
		如有一个 untyped 跟一个 非untyped 运算时
		会转换到 非untyped 的类型 
	*/

}
