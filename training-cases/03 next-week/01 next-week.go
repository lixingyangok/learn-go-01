// +build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	// ▼制造一个数组，有7个成员，分别是下周的7天
	// var slice01 = [7]string
	var firtTime = time.Now().Add(time.Hour * 24 * -19)
	var nowWeekday = int(firtTime.Weekday())
	var theGapToNextMonday = 7 - nowWeekday + 1 //距离周一还有几天
	var nextWeek = []string{}                   //下周
	for idx := 0; idx < 7; idx++ {
		thisTime := firtTime.Add(time.Hour * 24 * time.Duration(theGapToNextMonday+idx))
		nextWeek = append(
			nextWeek,
			thisTime.Format("\n2006/01/02 Mon"),
		)
	}
	fmt.Println("下周七天：", nextWeek)
}

