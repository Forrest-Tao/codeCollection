package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		}()
	}
	wg.Wait()
	//load store swap compareAndSwap
	//atomic.SwapInt64()
}

var lock uint32

func spinLock() {
	//尝试加锁，将无锁状态转为有锁
	for !atomic.CompareAndSwapUint32(&lock, 0, 1) {
		//加锁失败
		//自旋等待
	}
	//临界区域
	//handler....
	//释放锁
	atomic.StoreUint32(&lock, 0)
}
