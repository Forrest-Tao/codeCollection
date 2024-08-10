package main

import (
	"fmt"
	"time"
)

type T int

var (
	ch    chan T
	value T
	hook1 int
	hook2 int
)

func init() {
	ch = make(chan T, 10)
	value = 0
	hook1 = 0
	hook2 = 0
}

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func ReaderA() {
	for v := range ch {
		fmt.Println("in ReadA ", v)
		if hook1 == 5 {
			close(ch)
			fmt.Println("Closed in A")
			break
		}
		hook1++
	}
}

func ReaderB() {
	for v := range ch {
		fmt.Println("in ReadB ", v)
		if hook2 == 2 {
			close(ch)
			fmt.Println("Closed in B")
			break
		}
		hook2++
	}
}

func WriterC() {
	for {
		if !IsClosed(ch) {
			ch <- value
		} else {
			fmt.Println("ch was closed with value", value)
			break
		}
		value++
	}
}

func main() {
	go ReaderA()
	go ReaderB()
	go WriterC()

	time.Sleep(time.Second * 5)
}
