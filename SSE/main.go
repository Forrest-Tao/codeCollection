package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// SSEHandler 处理 SSE 请求
func SSEHandler(c *gin.Context) {
	println("in")
	// 设置必要的响应头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// 定义一个通道，用于发送事件
	messageChan := make(chan string)

	// 启动一个 goroutine，定期发送事件
	go func() {
		for {
			select {
			case <-c.Request.Context().Done():
				fmt.Println("Client disconnected")
				close(messageChan)
				return
			case <-time.After(2 * time.Second):
				messageChan <- fmt.Sprintf("data: %s\n\n", time.Now().String())
			}
		}
	}()

	// 发送事件
	for msg := range messageChan {
		fmt.Fprintf(c.Writer, msg)
		c.Writer.Flush()
	}
	println("out	")
}
func demo(c *gin.Context) {
	c.String(200, "Hello World!")
}

func main() {
	router := gin.Default()

	router.GET("/events", SSEHandler)
	router.GET("/demo", demo)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
