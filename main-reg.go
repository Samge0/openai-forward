package main

import (
	"fmt"
	"net/http"
)

// handlerReg 处理注册的地址转发
func handlerReg(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "https://platform.openai.com")
}

func main() {
	http.HandleFunc("/", handlerReg)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
