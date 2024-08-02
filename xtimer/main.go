package xtimer

import (
	"context"
	"sync"
	"time"
)

type XTimer struct {
	_m        sync.Mutex
	_timer    *time.Timer
	_stopChan chan struct{}
	_stop     bool //判断是否停止
}

func (t *XTimer) stop() error {
	t._m.Lock()
	defer t._m.Unlock()
	t._stopChan <- struct{}{}
	t._stop = true
	return nil
}

func StartBlockingTimer(duration time.Duration, fn func()) (*time.Timer, chan struct{}) {
	//第一次执行，没有等待时间，直接执行
	var _timer *time.Timer = time.NewTimer(0)
	_stopChan := make(chan struct{})
	go func() {
	OUT:
		for {
			select {
			case <-_timer.C:
				fn()
				_timer.Reset(duration)
			case <-_stopChan:
				break OUT
			}
		}
	}()
	return _timer, _stopChan
}

func XTicker(ctx context.Context, duration time.Duration, delay bool, fn func()) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	if !delay {
		fn()
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fn()
		}
	}
}
