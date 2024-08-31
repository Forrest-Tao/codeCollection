package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 验证ctx的共享性质
func do1(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("in do1")
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Duration(500) * time.Millisecond):
				fmt.Println("------do1-----")
			case <-ctx.Done():
				fmt.Println("do1 done")
				return
			}
		}
	}()
}

// do1和do2 同一个ctx
func do2(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("in do2")
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Duration(1) * time.Second):
				fmt.Println("------do2-----")
			case <-ctx.Done():
				fmt.Println("do2 done")
				return
			}
		}
	}()
}

// do3的ctx是do1的子ctx
func do3(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("in do3")
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Duration(300) * time.Millisecond):
				fmt.Println("------do3-----")
			case <-ctx.Done():
				fmt.Println("do3 done")
				return
			}
		}
	}()
}
