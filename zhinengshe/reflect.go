/*
 * @Author: 李星阳
 * @Date: 2020-08-28 20:19:38
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-08-28 20:29:35
 * @Description:
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 1.5
	var b = int(a)
	fmt.Printf("%v\n", reflect.TypeOf(a))
	fmt.Printf("%v\n", reflect.TypeOf(b))
}
