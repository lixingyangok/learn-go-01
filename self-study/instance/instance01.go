/*
 * @Author: 李星阳
 * @Date: 2021-03-03 11:26:24
 * @LastEditors: 李星阳
 * @LastEditTime: 2021-03-03 11:55:50
 * @Description:
 */
package instance

import "fmt"

/*
	Go语言的“基本”类型有（注意是“基本”类型
	bool
	string
	int、int8、int16、int32、int64
	uint、uint8、uint16、uint32、uint64、uintptr
	byte // uint8 的别名
	rune // int32 的别名 代表一个 Unicode 码
	float32、float64
	complex64、complex128
*/

func F01() {
	// var 变量名 类型 = 值
	var i01 int       // 声明不赋值（默认是零值）
	var i02 int = 100 // 声明赋值
	var i03, i04 int  // 声明多个变量为同一类型
	// ▼打印
	fmt.Println("整形值：", i01, i02, i03, i04)
}

func F02() {
	var ( // 批量声明此
		f01 float64
		f02 float64 = 3.14
		f03 float64 = 100.0
	)
	// ▼短声明
	f11, f12 := 1.1, 2.3
	// ▼打印
	fmt.Println("浮点值：", f01, f02, f03)
	fmt.Println("浮点值：", f11, f12)
}
