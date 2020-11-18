// +build ignore

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ▼如果切片成员类型是 interface{} 可以存任意类型
	var slice01 = []interface{}{
		"字符ABC123", 123, []int{4,5,6},
	}
	// fmt.Println(slice01) // 打印
	// val, ok := slice01[0].(string) //查询某值是否为string
	// val2, ok2 := slice01[0].(int) //是否为int
	// fmt.Println(val, ok)
	// fmt.Println(val2, ok2)

	switch slice01[2].(type) {
	case []int:
		fmt.Println("这是一个整形切片") //会打印出来
	default:
		fmt.Println("类型未知")
	}

	//另外如果只是单纯的想知道变量的类型，可以使用reflect.TypeOf()
	val01 := "abc123"
	val02 := 1
	val03 := 1.0
	val04 := []int{1,2,3}
	fmt.Println(reflect.TypeOf(val01))	//打印结果：string
	fmt.Println(reflect.TypeOf(val02))	//打印结果：int
	fmt.Println(reflect.TypeOf(val03))	//打印结果：float64
	fmt.Println(reflect.TypeOf(val04))	//打印结果：[]int

	typeResult:= reflect.TypeOf(
		reflect.TypeOf("abc"),
	)
	fmt.Println(typeResult)	// *reflect.rtype
}

//例子
// var s interface{}
// switch s.(type) {
// case string:
// 		fmt.Println("这是一个string类型的变量")
// case int64:
// 		fmt.Println("这是一个你int64类型的变量")
// default:
// 		fmt.Println("以上类型都不是")
// }
