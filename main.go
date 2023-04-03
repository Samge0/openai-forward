package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	targetURL := "https://api.openai.com" // 替换为要转发的域名
	req, _ := http.NewRequest(r.Method, targetURL, r.Body)
	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to target URL: %v", err)
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
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
