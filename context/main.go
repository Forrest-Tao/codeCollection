package main

import (
	"context"
	"fmt"
)

// withAsyncCancel 装饰器函数，用来封装 goroutine 逻辑并处理错误和取消操作
func withAsyncCancel(_ context.Context, cancel context.CancelFunc, fn func() error) func() {
	return func() {
		go func() {
			// 确保 goroutine 中的 panic 不会导致程序崩溃
			defer func() {
				if r := recover(); r != nil {
					cancel() // 取消操作
				}
			}()

			// 执行目标函数
			if err := fn(); err != nil {
				cancel() // 取消操作
			}
		}()
	}
}

func main() {
	//能重复cancel
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	cancel()
	cancel()
	cancel()
	fmt.Println("----")
}
