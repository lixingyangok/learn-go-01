// +build ignore

package main

import (
	"fmt"
	"time"
)

/*
	打印内容：
	今天时间，今天时间格式化，今天星期
	明天时间，
	下周一时间
	下周七天的切片
*/

var arr = [...]int{1, 7, 6, 5, 4, 3, 2}

func main() {
	// 今天
	nowTime := time.Now()
	nowTimeFormat := nowTime.Format("2006.01.02 15:04:05.000 ★ Mon")
	nowTimeDay := int(nowTime.Weekday())
	fmt.Printf("现在时间：%+v\n", nowTime)
	fmt.Printf("现在时间：%+v\n", nowTimeFormat)
	fmt.Printf("现在星期：%+v\n", nowTimeDay)
	for idx := 1; idx <= 6; idx++ {
		newDay := nowTime.AddDate(0, 0, idx)
		newWeekDay := newDay.Weekday()
		fmt.Printf(
			"+%d天 = 周%v，%d天后是下周一\n",
			idx, int(newWeekDay), arr[newWeekDay],
		)
	}

	// 下周一
	nextMonday := nowTime.AddDate(0, 0, arr[nowTimeDay])
	fmt.Printf("下周一：%+v\n", nextMonday)
}
