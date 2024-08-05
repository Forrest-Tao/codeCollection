package spinLock

import (
	"runtime"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	flag int32
}

func (sl *SpinLock) Lock() {
	spin := 0
	for !atomic.CompareAndSwapInt32(&sl.flag, 0, 1) {
		spin++
		if spin%100 == 0 {
			time.Sleep(time.Millisecond)
		} else {
			runtime.Gosched()
		}
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreInt32(&sl.flag, 0)
}
