package main

import (
	"fmt"
	"sync"
)

func PrintEven(wg *sync.WaitGroup, even chan struct{}, old chan struct{}) {
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i += 2 {
			<-even
			fmt.Println(i)
			old <- struct{}{}
		}
	}()
}

func PrintOld(wg *sync.WaitGroup, even chan struct{}, old chan struct{}) {
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			<-old
			fmt.Println(i)
			even <- struct{}{}
		}
	}()
}
