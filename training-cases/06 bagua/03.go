// +build ignore

package main

import ("fmt")

func getSymbol (num int) string {
	if num==0 {
		return "■■■ ■■■"
	}
	return "■■■■■■■"
}

func main() {
	for i1:= 0; i1 <= 1; i1++ {
		for i2:= 0; i2 <= 1; i2++ {
			for i3:= 0; i3 <= 1; i3++ {
				fmt.Printf(
					"%s\n%s\n%s\n--------------\n",
					getSymbol(i1), getSymbol(i2), getSymbol(i3),
				)
			}
		}
	}
}



