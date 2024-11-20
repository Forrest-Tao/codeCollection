package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/runtime"
	"net/http"
	"os"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源连接，可以根据需要调整
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server started on :9091")
	runtime.Must(http.ListenAndServe(":9091", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()
	file, err := os.Open("lines.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if conn.WriteMessage(websocket.TextMessage, scanner.Bytes()) != nil {
			fmt.Println("Failed to send message:", err)
			break
		}
		time.Sleep(time.Second)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	conn.WriteMessage(websocket.CloseMessage, []byte(""))
}
