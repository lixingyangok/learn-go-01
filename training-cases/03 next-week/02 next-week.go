// +build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	var firstTime = time.Now().Add(time.Hour * 24 * 5)
	var nowWeekday = int(firstTime.Weekday())
	var gapToNextMonday = 8 - nowWeekday
	fmt.Println("初始时间：", firstTime)
	fmt.Println("初始时间是星期：", nowWeekday)

	var nextWeekSlice = []string{}
	for idx := 0; idx < 7; idx++ {
		thisTime := firstTime.Add(time.Hour * 24 * time.Duration(idx + gapToNextMonday))
		nextWeekSlice = append(nextWeekSlice, "\n" + thisTime.Format("2006.01.02  15:04:05.000"))
	}
	fmt.Println("下一周：", nextWeekSlice)
}

