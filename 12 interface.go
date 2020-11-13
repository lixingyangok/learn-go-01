// +build ignore

package main

import (
	"fmt"
)


func main() {
	// ▼如果切片成员类型是 interface{} 可以存任意类型
	var slice01 = []interface{}{
		"字符ABC123", 123, []int{4,5,6},
	}

	fmt.Println(slice01)
	val, ok := slice01[0].(string)
	val2, ok2 := slice01[0].(int)
	fmt.Println(val, ok)
	fmt.Println(val2, ok2)
}


//例子
// var s interface{}
// switch s.(type) {
// case string:
// 		fmt.println("这是一个string类型的变量")
// case int64:
// 		fmt.println("这是一个你int64类型的变量")
// default:
// 		fmt.println("以上类型都不是")
// }

// //另外如果只是单纯的想知道变量的类型，可以使用reflect.typeof()
// val := "abcdefg123"
// fmt.println(reflect.typeof(val))		//打印结果：string
// beego.Debug(reflect.typeof(val))	//Debug打印结果：string