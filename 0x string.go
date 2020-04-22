package main

import (
	"fmt"
	"strings"
)


func main() {
	var s01 = '字' //单引号只能包裹一个字符
	var s02 = "字符串"
	var s03 = `"多行"字符串`; //多行字符声明
	fmt.Println(s01)
	fmt.Println(s02)
	fmt.Println(s03)
	fmt.Println("-------------------------")
	fmt.Println(`字符串长度：`, len(s02))
	fmt.Println(`字符串拼接：`, s02 + s02)
	fmt.Println(fmt.Sprintf(`字符串模板： 欢迎“%s”今天第“%s”次登录`, "小明", "3"))
	fmt.Println(`字符串分割：`, strings.Split("分,割", ",")) //分割为数组
	fmt.Println(`字符串包含：`, strings.Contains(s02, "字"))
	fmt.Println(`字符串开头：`, strings.HasPrefix(s02, "字"))
	fmt.Println(`字符串结尾：`, strings.HasSuffix(s02, "串"))
	fmt.Println(`字符串中某字索引-正序：`, strings.Index(s03, `"`))
	fmt.Println(`字符串中某字索引-倒序：`, strings.LastIndex(s03, `"`))
	fmt.Println(`将数组成员连成字符串：`, strings.Join(strings.Split("分,割", ","), "+"))
	// fmt.Println(`将数组成员连成字符串：`, strings.Join(["11", "22", "33"], "+"))
	fmt.Println("-------------------------")
	for _, cur := range s02 {
		// fmt.Println(cur) //打印一堆数字
		fmt.Printf("%c, ", cur) //能正常打印 byte 和 rune 类型
	}
}