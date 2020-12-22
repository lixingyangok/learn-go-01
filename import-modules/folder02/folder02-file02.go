/*
 * @Author: 李星阳
 * @Date: 2020-11-18 15:12:42
 * @LastEditors: 李星阳
 * @LastEditTime: 2020-11-18 17:01:25
 * @Description:
 */

package badpackagename

// 一个包可由多个不同文件组成（同级文件）

// GetBiggest 取得最大数
func GetBiggest(numbers ...int) (result int) {
	for _, curNumber := range numbers {
		if curNumber > result {
			result = curNumber
		}
	}
	return
}

// GetBiggest02 取得最大数
func GetBiggest02(theSlice []interface{}) (result float64) {
	for _, cur := range theSlice {
		value, ok1 := cur.(int)
		if ok1 && float64(value) > result {
			result = float64(value)
		}
		value2, ok2 := cur.(float64)
		if ok2 && value2 > result {
			result = value2
		}
	}
	return
}
