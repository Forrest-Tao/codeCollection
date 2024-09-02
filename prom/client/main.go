package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	for i := 0; i < 2000; i++ {
		go func() {
			resp, err := client.Get("http://localhost:8080/api")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			resp.Body.Close()
		}()
		time.Sleep(10 * time.Millisecond) // 模拟请求间隔
	}
	time.Sleep(5 * time.Second) // 等待所有请求完成
}
