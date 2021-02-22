package main

import "fmt"

// ▼双击 build.cmd 可将此文件编译为 .wasm 
func main() {
	fmt.Println("这是从wasm打印的一段文字 01");
	fmt.Println("这是从wasm打印的一段文字 02");
}


// import "syscall/js"

// go env -w GOOS=js 
// go env -w GOARCH=wasm 
// import "syscall/js"

// func main() {
// 	toLog := js.Global().get("console").get("log")
// 	toLog.Invoke("123")
// 	js.Global().Call("eval", `
// 		console.log("456")
// 	`)
// }

// func main() {
//     f_fib := func(params []js.Value) {
//         var n = params[0].Int()  // 输入参数
//         var callback = params[1] // 回调参数
//         var result = fib(n)
//         // 调用回调函数，传入计算结果
//         callback.Invoke(result)
//     }
//     // 注册全局函数
//     js.Global().Set("fib", js.NewCallback(f_fib))
//     // 保持 main 函数持续运行
//     select {}
// }

// // 计算斐波那契数
// func fib(n int) int {
//     if n <= 0 {
//         return 0
//     }
//     var result = make([]int, n+1)
//     result[0] = 0
//     result[1] = 1
//     if n <= 1 {
//         return result[n]
//     }
//     for i := 2; i <= n; i++ {
//         result[i] = result[i-2] + result[i-1]
//     }
//     return result[n]
// }
