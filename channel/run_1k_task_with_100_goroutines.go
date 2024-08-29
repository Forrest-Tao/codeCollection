package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var cnt atomic.Int64

func worker(workerId int, wg *sync.WaitGroup, ch chan func(taskId int64, workerId int)) {
	defer wg.Done()
	for f := range ch {
		cnt.Add(1)
		f(cnt.Load(), workerId)
	}
}

func run() {
	var (
		wg   sync.WaitGroup
		todo = make(chan func(taskId int64, workerId int), 100)
	)

	for workerId := range 100 {
		wg.Add(1)
		go worker(workerId, &wg, todo)
	}

	f := func(taskId int64, workerId int) {
		time.Sleep(time.Millisecond * time.Duration(100))
		fmt.Printf("---worker: %d is handling %d---\n", workerId, taskId)
	}

	for range 1000 {
		todo <- f
	}

	close(todo)
	wg.Wait()
}
