// +build ignore123


package main

import (
	"fmt"
	"syscall/js"
)

// 通过递归计算第 i 个斐波那契数
func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(fib(args[0].Int()))
}

// function definition
func add(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(i[0].Int() + i[1].Int())
}

func getBigNumber(n01 js.Value, args []js.Value) interface{}{
	iTimes := args[0].Int()
	iResult := 0
	for idx := 0; idx < iTimes; idx++ {
		iResult += idx
	}
	return js.ValueOf(iResult);
}

func main() {
	fmt.Println("这是从wasm打印的一段文字 go-01");
	/* 
		alert := js.Global().Get("alert") // 获取全局的 alert 对象
		alert.Invoke("Hello World!") // 等价于在 js 中调用 window.alert("Hello World")
	*/

	// exposing to JS
	// 注册 go 的方法到 js 的 window 对象
	js.Global().Set("addNumber", js.FuncOf(add))
	js.Global().Set("getBigNumber", js.FuncOf(getBigNumber))

	done := make(chan int, 0)
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<- done
}


