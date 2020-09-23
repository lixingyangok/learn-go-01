// +build ignore

package main

import (
	"fmt"
	"time"
)

func getNextWeek(oneTime time.Time) []string {
	fmt.Println("日期", oneTime)
	var gapToNextWeek = 8 - int(oneTime.Weekday())
	var result = []string{}
	for idx := 0; idx < 7; idx++ {
		thisDay := oneTime.Add(time.Hour * 24 * time.Duration(gapToNextWeek + idx))
		result = append(
			result, thisDay.Format("\n2006.01.02  15:04:05  Mon"),
		)
	}
	return result
}

func main() {
	var result01 = getNextWeek(time.Now().Add(time.Hour * 24 * 3))
	var result02 = getNextWeek(time.Now().Add(time.Hour * 24 * 36))
	fmt.Println("result01", result01)
	fmt.Println("result02", result02)
}
