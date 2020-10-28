// +build ignore

package main

import (
	"fmt"
)

func numsToStr(num... int) string {
	result := ""
	for _, cur := range num {
		if cur == 0 {
			result += "○"
		}else if cur == 1 {
			result += "●"
		}
	}
	return result
}

func main() {
	for i01 := 0; i01 <= 1; i01++{
		for i02 := 0; i02 <= 1; i02++{
			for i03 := 0; i03 <= 1; i03++{
				for i04 := 0; i04 <= 1; i04++{
					for i05 := 0; i05 <= 1; i05++{
						for i06 := 0; i06 <= 1; i06++{
							for i07 := 0; i07 <= 1; i07++{
								for i08 := 0; i08 <= 1; i08++{
									thisSymbol := numsToStr(i01,i02,i03,i04,i05,i06,i07,i08)
									fmt.Println(thisSymbol)
								}
							}
						}
					}
				}
			}
		}
	}
}
