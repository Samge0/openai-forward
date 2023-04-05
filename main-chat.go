package main

import (
	"fmt"
	"net/http"
)

// handlerChat 处理聊天界面的地址转发
func handlerChat(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "https://chat.openai.com")
}

func main() {
	http.HandleFunc("/", handlerChat)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
