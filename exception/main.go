package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("in main", r)
		}
	}()
	go callFunc()
	cnt := 0
	for cnt < 5 {
		time.Sleep(time.Second)
		println("----")
		cnt++
	}
}

func callFunc() {
	//自己的panic自己recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("in callFunc", r)
		}
	}()
	panic("Something went wrong")
}
