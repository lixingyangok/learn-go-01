// 这里练习数据类型
package main
import (
	"fmt"
	"reflect"
)
const sLine = "\n\n------------------------------------"

func main() {
	fmt.Println("\n\n\n\n■■ 开始 ■■■■■■■■■■■■■■■■■■■■■■■■■■")
	fmt.Println(
		"\n 0:", 		reflect.TypeOf(0), //0被视为整形
		"\n 1000:", 	reflect.TypeOf(1000),
		"\n 0.11:", 	reflect.TypeOf(0.11),
		"\n 0.99:", 	reflect.TypeOf(0.99),
		sLine,
	)
	fmt.Println(
		"\n", 'A', reflect.TypeOf('A'), //0被视为整形
		"\n", "A", reflect.TypeOf("A"),
		sLine,
	)

	var int01 = 0 //默认int
	// int01 = 0.5 //整形不可被赋值浮点型
	var int02 int32 = 100
	var int03 int64 = 100
	var int01Type = reflect.TypeOf(int01).String() //将类型转化为字符
	fmt.Println(
		int01Type, reflect.TypeOf(int02), reflect.TypeOf(int03),
		sLine,
	)

	var float01 = 0.0 //默认 float64
	float01 = 1 //浮点可以被赋值整形
	var float02 float32 = 0;
	var float03 float64 = 0;


	fmt.Println(
		reflect.TypeOf(float01), reflect.TypeOf(float02), reflect.TypeOf(float03),
		sLine,
	)


	// var int21 = 0.5;
	// // var int02 = int64{0}
	// fmt.Println(reflect.TypeOf(int12))
	// fmt.Println(reflect.TypeOf(int21))
	// // fmt.Println(reflect.TypeOf(int02))
	// fmt.Println(reflect.TypeOf('福'))
	// fmt.Println(reflect.TypeOf("A"))
}
