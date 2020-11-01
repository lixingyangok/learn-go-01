package test

import (
	"fmt"
	"testing"
)

/*
	1、文件名须是 xxx_test.go 格式，即以 _test.go 结尾。例如 hello_test.go
	2、方法名须是 Test 打头，并且形参为 (t *testing.T)。
		例如：func Test_fn01(t *testing.T) {}
		例如：func TestHelloWorld(t *testing.T) {}
*/

// ▼测试方法01
func Test_fn01(t *testing.T) {
	// 相同目录的多个测试文件中，不能存在相同名称的函数名，即测试方法禁止重名
	aa := 3
	bb := 7
	fmt.Println("测试01-开始：●●●●●●●●●●●")
	fmt.Println("aa * bb =", aa*bb)
	fmt.Println("测试01-结束：●●●●●●●●●●●")
	t.Log("测试01-结束：●●●●●●●●●●●") //这个偶尔没打印不出来
}

// ▼测试方法02
func Test_fn02(t *testing.T) {
	// 相同目录的多个测试文件中，不能存在相同名称的函数名，即测试方法禁止重名
	aa := 4
	bb := 8
	fmt.Println("测试02-开始：●●●●●●●●●●●")
	fmt.Println("aa * bb =", aa*bb)
	fmt.Println("测试02-结束：●●●●●●●●●●●")
}

/* 
	如何执行【文件内所有方法】？
	先 cd 到当前目标
	执行：go test 01_test.go -v
	go test 01_test.go -v Test_fn01

	▼以下是有效方法，注，此处没输入文件名，只输入方法名
	go test -run "^Test_fn01$" -v
	go test -run "^Test_fn01$"
	go test -run "Test_fn01"
	go test -run Test_fn01
*/

/*
	如何执行具体的【方法】？
	go test 01_test.go -v Test_fn01
	
	▼以下是有效方法，注，此处没输入文件名，只输入方法名
	有效：go test -run "^Test_fn01$" -v
	有效：go test -run "^Test_fn01$"
	有效：go test -run "Test_fn01"
	有效：go test -run Test_fn01
*/
