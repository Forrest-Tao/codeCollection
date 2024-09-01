package main

import (
	"fmt"
)

func multi_choice() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	// 向每个 channel 写入数据
	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	// 使用 select 从多个已准备好的 channel 中选择
	select {
	case val := <-ch1:
		fmt.Println("Received from ch1:", val)
	case val := <-ch2:
		fmt.Println("Received from ch2:", val)
	case val := <-ch3:
		fmt.Println("Received from ch3:", val)
	}
}
