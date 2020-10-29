package main

import (
	"fmt"
	"testing"
)
/* 
1、文件名须以"_test.go"结尾。例如 /hello_test.go
2、方法名须以"Test"打头，并且形参为 (t *testing.T)
*/

/* 
先 cd 到当前目标
执行：【go test -v 01_test.go】可执行当前文件中的所有函数
如执行某特定函数：go test -v 依赖文件 -test.run 方法名。例如：go test -v 01_test.go -test.run Test_fn01
▲上述执行未成功

有效：go test -run "^Test_fn01$" -v
有效：go test -run "^Test_fn01$"
有效：go test -run "Test_fn01"
有效：go test -run Test_fn01
*/


func Test_fn01(t *testing.T) {
	// 相同目录的多个文件中不能存在相同名称的函数名，即测试方法禁止重名
	fmt.Println("test")
	fmt.Println("test")
	fmt.Println("test")
}