package main

import (
	"log"
	"net/http"
)

/*
	这个 go 文件编译后得到的 exe 文件可以启动一个服务，
	访问这个服务，可访问同级目录的 index.html
	访问 index 用于调试
*/

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8000", mux))
}
