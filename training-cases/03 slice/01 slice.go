// +build ignore

package main
import (
	"fmt"
	"time"
)

func main() {
	// ▼制造一个数组，有7个成员，分别是下周的7天
	// var slice01 = [7]string
	var nowTime = time.Now();
	var nowWeekday = int(nowTime.Weekday());
	fmt.Println(nowTime)
	fmt.Println("今天星期：", nowWeekday)
	fmt.Println(1, "天后：", time.Now().Add(time.Hour * 24 ))

}
