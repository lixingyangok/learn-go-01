/*
 * @Author: 李星阳
 * @Date: 2020-04-23 13:20:43
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-05-08 13:59:33
 * @Description: 
 */
package main

import "fmt"

func main() {
	// var age = 17
	// ▼将age声明到这个判断作用域内
	if age := 10 + 7; age >= 18 {
		fmt.Println(age, "开始赌博------")
	} else if age > 10 {
		fmt.Println(age, "即将赌博---")
	} else {
		fmt.Println(age, "不能赌博-")
	}

	var iFinger = 4
	switch iFinger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	default:
		fmt.Println("其它")
	}

	switch iNum:=3; iNum {
	case 1, 3, 5:
		fmt.Println("加班")
	case 2, 4, 6, 7:
		fmt.Println("休息")
	}
	
	iNum:=20
	switch {
	case iNum > 0 && iNum <= 80:
		fmt.Println("常人")
		fallthrough //再执行一行
	case iNum <= 100:
		fmt.Println("高寿")
	case iNum >= 100:
		fmt.Println("成仙")
	}
}
