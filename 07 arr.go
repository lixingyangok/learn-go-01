package main
import "fmt"

var a01 [3]int//声明时要指定：长度、成员类型
var a02 [3]bool //如果没有初始化，成员默认为零值：false, 0, "", 
var a03 [2]{true, false } //第3个成员是false
func main() {
						fmt.Print(a01)
}