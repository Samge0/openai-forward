package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	targetURL := "https://api.openai.com" // 替换为要转发的域名
	modifiedRequest, _ := http.NewRequest(http.MethodPost, targetURL, r.Body)
	modifiedRequest.Header = r.Header
	resp, err := client.Do(modifiedRequest)
	if err != nil {
		log.Fatalf("Error sending request to target URL: %v", err)
	}
	defer resp.Body.Close()

	w.Header().Add("Content-Type", "application/octet-stream")
	for {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		if len(body) == 0 {
			break
		}

		_, err = w.Write(body)
		if err != nil {
			log.Fatalf("Error writing response to client: %v", err)
		}

		w.(http.Flusher).Flush()
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
