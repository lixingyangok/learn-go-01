/*
 * @Author: 李星阳
 * @Date: 2020-05-09 10:48:19
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-05-09 13:46:32
 * @Description: 
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var aa = 3 / 2 // =1, int型
	fmt.Println("3/2 =", aa, reflect.TypeOf(aa), "\n")  // =1，int型

	var bb = float32(3) / 2 //float32型
	fmt.Println("float32(3) / 2 = ", bb, "\n")


	fmt.Println("3/1.5 =", 3 / 1.5, reflect.TypeOf(3/1.5)) // =2，float64型
	fmt.Println("3/1.0 =", 3 / 1.0, reflect.TypeOf(3/1.0)) // =3，float64型
	fmt.Println(
		reflect.TypeOf(aa), "\n",
		reflect.TypeOf(bb),
	)
}
