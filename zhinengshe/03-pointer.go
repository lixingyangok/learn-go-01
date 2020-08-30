package main

import "fmt"

func main() {
	var aa = 5
	var bb = &aa
	fmt.Printf("%v %v \n", aa, *bb)
	*bb = 6
	fmt.Printf("%v %v \n", aa, *bb)
}
