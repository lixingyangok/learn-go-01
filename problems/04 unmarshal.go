// +build ignore

package main

import (
	"fmt"
	"encoding/json"
)

type Animal struct {
    Name  string
    Order string
}

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	// {"Name": "好的",    "Order": "Dasyuromorphia"}

	var animals []Animal
	
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}