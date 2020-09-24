// +build ignore

package main
import (
	"fmt"
)

func main() {
	for idx := 1; idx<=9; idx ++ {
		for jj := 1; jj<=idx; jj ++ {
			fmt.Print(jj, "*", idx, "=", idx*jj, "  ")
		}
		fmt.Println()
	}
}
