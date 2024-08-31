package main

import (
	"context"
	"fmt"
	"sync"
	"time"
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

	run()
}

func handle(wg *sync.WaitGroup, ch chan string, ctx context.Context, cancelFunc context.CancelFunc) {
	tag := 0
	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-ctx.Done():
				cancelFunc()
				tag = 1
			}
			fmt.Println("out")
			if tag == 1 {
				break
			}
		}
	}()
}

func run() {
	var (
		wg  sync.WaitGroup
		ctx = context.Background()
	)

	newctx, cancel := context.WithCancel(ctx)
	ch := make(chan string, 1)
	wg.Add(1)
	handle(&wg, ch, newctx, cancel)
	for i := range 3 {
		ch <- fmt.Sprintf("--%d--\n", i)
	}
	cancel()
	wg.Wait()
}

func runShare() {
	var wg sync.WaitGroup
	wg.Add(3)

	parent, cancel1 := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	do1(parent, &wg)
	do2(parent, &wg)
	do3(child, &wg)

	time.Sleep(time.Second * 3)
	fmt.Println("after 3 s,cancel func")
	cancel1()
	wg.Wait()
}
