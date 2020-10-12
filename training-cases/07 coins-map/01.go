// +build ignore

package main

import (
	"fmt"
)

func main() {
	// var slice01 = make([]string, 0)
	var flag = 1
	for i01 := 0; i01 <= 1; i01++ {
		for i02 := 0; i02 <= 1; i02++ {
			for i03 := 0; i03 <= 1; i03++ {
				for i04 := 0; i04 <= 1; i04++ {
					for i05 := 0; i05 <= 1; i05++ {
						for i06 := 0; i06 <= 1; i06++ {
							// fmt.Println(flag, "：", i01,i02,i03,i04,i05,i06)
							var aa, err = fmt.Printf("%v：%v%v%v%v%v%v\n", flag, i01, i02, i03, i04, i05, i06)
							fmt.Println(aa, err)
							// slice01 = append(slice01, )
							flag++
						}
					}
				}
			}
		}
	}
}
