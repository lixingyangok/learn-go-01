package main

import (
	"fmt"
	"reflect"
	"strings"
)
var pt = fmt.Println

func main() {
	var s01 = "磨穿铁砚-天道酬勤-"
	var s02 = "磨穿铁砚-天道酬勤-"
	var splitStr = strings.Split(s01, "-")
	pt("字节长度：", len(s01))
	pt("字符长度：", len([]rune(s01)))
	pt("拼接字符：", s01 + "----" + s02)
	
	pt("分割字符串为切片（按某字符）：", splitStr)
	pt("分割之后的类型是：", reflect.TypeOf(splitStr)) //切片
	pt("查看是否包含某字符：", strings.Contains(s01, "铁"))
	pt("查看是否某字开头：", strings.HasPrefix(s01, "磨"))
	pt("查看是否某字结尾：", strings.HasSuffix(s01, "勤"))
	pt("查看某字所在索引：", strings.Index(s01, "-"))
	pt("查看某字所在索引（倒序：", strings.LastIndex(s01, "-"))
}