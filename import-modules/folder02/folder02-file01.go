/*
 * @Author: 李星阳
 * @Date: 2020-11-18 15:12:42
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-11-18 15:36:25
 * @Description:
 */

package badpackagename

// 此处包命名应与包所在目录同名（这里不同名，不应这样做)
// 另外提示中有表示：包名不应大小写混用，应全小写，或全大写

// GetSumFn 一求和函数
func GetSumFn(numbers ...int) (result int) {
	for _, curNumber := range numbers {
		result += curNumber
	}
	return
}
