package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// 处理api的地址转发
func handlerApi(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "https://api.openai.com")
}

// handlerReg 处理注册的地址转发
func handlerReg(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "https://platform.openai.com")
}

// handlerChat 处理聊天界面的地址转发
func handlerChat(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "https://chat.openai.com")
}

func handler(w http.ResponseWriter, r *http.Request, forwardUrl string) {
	client := &http.Client{}

	// 替换为要转发的域名
	urlPath := strings.Replace(r.RequestURI, "/forward-reg", "", 1)
	urlPath = strings.Replace(r.RequestURI, "/forward-chat", "", 1)
	targetURL := forwardUrl + urlPath
	req, _ := http.NewRequest(r.Method, targetURL, r.Body)
	req.Header = r.Header

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make request to target URL", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	for key, value := range resp.Header {
		w.Header().Set(key, value[0])
	}
	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Fatalf("Error copying response to client: %v", err)
	}
}
func main() {
	http.HandleFunc("/", handlerApi)
	http.HandleFunc("/forward-reg", handlerReg)
	http.HandleFunc("/forward-chat", handlerChat)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
