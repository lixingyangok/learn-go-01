// +build ignore

package main

import (
	// "fmt"
	"syscall/js"
)

func main() {
	// function definition
	func add(this js.Value, i []js.Value) interface{} {
		return js.ValueOf(i[0].Int() + i[1].Int())
	}
	// exposing to JS
	js.Global().Set("add", js.FuncOf(add))
}

