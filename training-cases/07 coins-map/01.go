// +build ignore

package main

import (
	"fmt"
)

func numToString(num int) (string){
	if num == 1 {
		return "●"
	}else if num == 0 {
		return "○"
	}
	return "-"
}

func main() {
	// var slice01 = make([]string, 0)
	var flag = 1
	for i01 := 0; i01 <= 1; i01++ {
		for i02 := 0; i02 <= 1; i02++ {
			for i03 := 0; i03 <= 1; i03++ {
				for i04 := 0; i04 <= 1; i04++ {
					for i05 := 0; i05 <= 1; i05++ {
						for i06 := 0; i06 <= 1; i06++ {
							var coinsList = numToString(i01) + numToString(i02) + numToString(i03) + numToString(i04) + numToString(i05) + numToString(i06)
							fmt.Println(flag, coinsList)
							flag++
						}
					}
				}
			}
		}
	}
}
