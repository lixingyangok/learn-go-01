// +build ignore

package main

import (
	"fmt"
	"reflect"
	"strings"
	// "strconv"
	// "utf8"
)

func main() {
	fmt.Println("\n")
	// ▼字符
	var s00 byte = 97                                         //byte是uint8的别名
	var s01 byte = 'a'                                        //单引号只能包裹一个字符
	fmt.Println("数字与byte对比：", s00 == s01, s01 == 97)          // true
	fmt.Println("byte类型就是uint8类型：", s00, reflect.TypeOf(s00)) //uint8
	fmt.Println("byte的type类型：", s01, reflect.TypeOf(s01))     //uint8
	// 分界===========================================================================
	// ▼字符串
	var s02 = "abc字符串"
	var s03 = `"多行"字符串` //多行字符声明
	var s04 = s02[0]    //
	var s05 = "These are English words"
	// ▼取下标
	fmt.Println("直接取下标：", s02[0])                 //得到数字，因为字符的背后是 []byte
	fmt.Println("取英文下标：", string(s02[1]))         //得到英文
	fmt.Println("取中文下标：", string([]rune(s02)[3])) //得到中文

	//
	fmt.Println(s02, reflect.TypeOf(s02))
	fmt.Println(s03, reflect.TypeOf(s03))
	fmt.Println(s04, reflect.TypeOf(s04))
	fmt.Println("-------------------------")
	fmt.Println(`字符串长度：`, len(s02), len([]rune(s02)))
	fmt.Println(`字符串大小写转换：`, strings.ToUpper(s05)) //ToLower
	// fmt.Println(`字符串长度：`, utf8.RuneCountInString(s02))
	fmt.Println(`字符串拼接：`, s02+s03)
	fmt.Println(fmt.Sprintf(`字符串模板： 欢迎“%s”今天第“%s”次登录`, "小明", "3"))
	fmt.Println(`字符串分割：`, strings.Split("分,割", ",")) //分割为数组
	fmt.Println(`字符串包含：`, strings.Contains(s02, "字"))
	fmt.Println(`字符串开头：`, strings.HasPrefix(s02, "字"))
	fmt.Println(`字符串结尾：`, strings.HasSuffix(s02, "串"))
	fmt.Println(`字符串中某字索引-正序：`, strings.Index(s03, `"`))
	fmt.Println(`字符串中某字索引-倒序：`, strings.LastIndex(s03, `"`))
	fmt.Println(`将数组成员连成字符串：`, strings.Join(strings.Split("分,割", ","), "+"))
	fmt.Println(`删除首尾空格`, strings.Trim("     	--abc123  - ", " -	"))
	// fmt.Println(`将数组成员连成字符串：`, strings.Join(["11", "22", "33"], "+"))
	fmt.Println("-------------------------")
	for idx, cur := range s02 {
		fmt.Print(idx, " - ", cur, " - ") //索引【不准】
		fmt.Printf("%c\n", cur)           //能正常打印 byte 和 rune 类型
	}
	fmt.Println("-------------------------")
	for idx, cur := range []rune(s02) {
		fmt.Println(idx, " - ", string(cur)) //索引【准】
	}
	fmt.Println("-------------------------")

	var str = "北京abc"
	oneRune := []rune(str)
	oneInt := oneRune[0]
	oneStr := string(oneInt)
	//
	fmt.Println(str, reflect.TypeOf(str))
	fmt.Println(oneRune, reflect.TypeOf(oneRune))
	fmt.Println(oneInt, reflect.TypeOf(oneInt))
	fmt.Println(oneStr, reflect.TypeOf(oneStr))
	// fmt.Print(str3, strtype)
	// for i := 0; i < len(str2); i++ {
	//     fmt.Println(str2[i])
	// }
}
