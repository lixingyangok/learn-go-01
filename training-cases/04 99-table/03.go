// +build ignore

// 2020.10.19 10:34:3 星期一

package main
import ("fmt")

func main() {
	for ii:= 1; ii<=9; ii++{
		for jj := 1; jj<=ii; jj++{
			fmt.Printf("%d*%d=%2d  ", jj, ii, jj * ii)
		}
		fmt.Println()
	}
}
