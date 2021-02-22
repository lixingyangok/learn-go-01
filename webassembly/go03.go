// +build ignore123
// 双击 build.cmd 可将此文件编译为 .wasm

package main

import (
	"fmt"
	"syscall/js"
)

func getBigNumber(n01 js.Value, args []js.Value) interface{}{
	iTimes := args[0].Int()
	iResult := 0
	for idx := 0; idx < iTimes; idx++ {
		iResult += idx
	}
	return js.ValueOf(iResult);
}

func getSum(n01 js.Value, args []js.Value)  interface{}{
	iTimes := args[0].Int()
	result := 0.0
	for idx := 0; idx < iTimes; idx++ {
		result += 0.001
	}
	return js.ValueOf(result);
}

func main() {
	fmt.Println("这是从wasm打印的一段文字 go03 - 3");
	done := make(chan int, 0)
	js.Global().Set("getBigNumber", js.FuncOf(getBigNumber))
	js.Global().Set("getSum", js.FuncOf(getSum))
	<- done
}


