// +build ignore

// 2020.10.19 10:34:3 星期一

package main
import ("fmt")

func main() {
	for idx := 1; idx <= 9; idx++ {
		for jj := 1; jj<=idx; jj++{
			result := jj* idx
			tail := "  "
			if result < 10 {
				tail += " "
			}
			fmt.Printf("%d*%d=%d%s", jj, idx, result, tail)
		}
		fmt.Println("")
	}
}
