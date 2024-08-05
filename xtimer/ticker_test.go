package xtimer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	now := time.Now()
	_timer := time.NewTimer(0)
	<-_timer.C
	fmt.Println("Timer triggered", "after", time.Since(now))
}

func TestMap(t *testing.T) {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}

func TestSyncMap(t *testing.T) {
	sync.Map{}
}
