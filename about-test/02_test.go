package test

import (
	"fmt"
	"testing"
)

// ▼测试方法（不用显示的调用）当执行 go test 时会自动执行它
func Test_A(t *testing.T) {
	fmt.Println("Test_A")
}

func Test_B(t *testing.T) {
	fmt.Println("Test_B")
}
