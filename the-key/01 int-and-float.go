package main

import "fmt"

func main() {
	var slice01 = [...]int{1,2,3}
	var aa = len(slice01)
	var bb = float32(aa) / 2
	fmt.Println(aa, bb)
}