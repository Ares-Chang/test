package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 定义一个处理静态文件的处理器
	fileServer := http.FileServer(http.Dir("./"))

	// 注册文件处理器，将根 URL ("/") 映射到该处理器
	http.Handle("/", fileServer)

	// 启动服务器并监听指定的端口
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
