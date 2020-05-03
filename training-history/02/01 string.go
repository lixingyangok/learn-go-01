/*
 * @Author: 李星阳
 * @Date: 2020-04-30 12:51:51
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-05-03 13:44:23
 * @Description: 
 */

package main
import(
	"fmt"
	"strings"
	"reflect"
)
var pt = fmt.Println

func main() {
	var s01 = "好好学习"
	var s02 = "天天向上"
	var s03 = s01 + "-" + s02 //连接字符串
	var s04 = `
		第一行
		第二行
	`
	pt("字【节】长度：", len(s01))
	pt("字【符】长度：", len([]rune(s02)))
	pt("连接多个字符串：", s03)
	pt("多行文本修剪首尾空格：", strings.TrimSpace(s04))
	// 
	var splitS03 = strings.Split(s03, "-") //转化为切片
	splitS03 = append(splitS03, "财务自由")
	pt("将字符分割成切片：", splitS03, reflect.TypeOf(splitS03))
	pt("查看是否包含某字符：", strings.Contains(s01, "学"))
	pt("查看是否某字开头/结尾：", strings.HasPrefix(s01, "学"), strings.HasSuffix(s01, "习"))
	pt("查看某字所在索引正序/倒序：", strings.Index(s01, "好"), strings.LastIndex(s01, "好"))
	pt("将数组成员连接成字符串：", strings.Join([]string{"吕布", "赵云"}, "--------"))
	// 遍历字符串
	for idx, val := range s01 {
		fmt.Println(idx, "-", val, "-", string(val))
	}
}

