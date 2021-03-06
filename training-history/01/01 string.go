/*
 * @Author: 李星阳
 * @Date: 2020-05-03 13:04:35
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-12-02 20:42:16
 * @Description:
 */

package main

import (
	"fmt"
	"reflect"
	"strings"
)

var pt = fmt.Println

func reverseStr(str string) (result string) {
	for _, cur := range str {
		result = string(cur) + result
	}
	return result
}
func reverseStr02(str string) string {
	var myRune = []rune(str)
	var length int = len(myRune)
	var result float32 = float32(length) / 2
	fmt.Println(length, result)
	for idx := 0; idx < length/2; idx++ {
		fmt.Println(idx)
		var lastIdx = length - 1 - idx
		myRune[idx], myRune[lastIdx] = myRune[lastIdx], myRune[idx]
	}
	return string(myRune)
}

func main() {
	var s01 = "磨穿铁砚-"
	var s02 = "天道酬勤-"
	var splitStr = strings.Split(s01, "-")
	pt("字节长度：", len(s01))
	pt("字符长度：", len([]rune(s01)))
	pt("拼接字符：", s01+"----"+s02)

	pt("分割字符串为切片（按某字符）：", splitStr)
	pt("分割之后的类型是：", reflect.TypeOf(splitStr)) //切片
	pt("查看是否包含某字符：", strings.Contains(s01, "铁"))
	pt("查看是否某字开头：", strings.HasPrefix(s01, "磨"))
	pt("查看是否某字结尾：", strings.HasSuffix(s01, "勤"))
	pt("查看某字所在索引：", strings.Index(s01, "-"))
	pt("查看某字所在索引（倒序：", strings.LastIndex(s01, "-"))
	for idx, val := range s01 {
		fmt.Println(idx, "-", val, "-", string(val))
	}
	// var numToStr01 = string(30952)              // 真·转文字【这行能执行，报错】
	var numToStr02 = fmt.Sprintf("%v", 123.456) // 数字转文字（字符型数字
	// pt("数字转文字：", numToStr01)
	pt("数字转文字：", numToStr02)

	pt("反转字符：", reverseStr("123"))
	pt("反转字符：", reverseStr02("123"))
}
