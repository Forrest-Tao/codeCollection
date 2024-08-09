package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	//close(ch) // 关闭 channel

	// 尝试从关闭的 channel 中接收数据
	value, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed.")
	} else {
		fmt.Printf("Received value: %d\n", value)
	}

	// 尝试再接收一次
	value, ok = <-ch
	if !ok {
		fmt.Println("Channel is closed.")
	} else {
		fmt.Printf("Received value: %d\n", value)
	}
}
