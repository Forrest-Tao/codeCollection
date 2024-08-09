package main

/*
	使用分段锁 减轻
*/

import (
	"fmt"
	"sync"
)

const numSegments = 32

type SafeMap struct {
	mu [numSegments]sync.RWMutex
	m  [numSegments]map[string]int
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{}
	for i := 0; i < numSegments; i++ {
		sm.m[i] = make(map[string]int)
	}
	return sm
}

func (sm *SafeMap) segment(key string) *sync.RWMutex {
	return &sm.mu[len(key)%numSegments]
}

func (sm *SafeMap) Set(key string, value int) {
	seg := sm.segment(key)
	seg.Lock()
	defer seg.Unlock()
	sm.m[len(key)%numSegments][key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	seg := sm.segment(key)
	seg.RLock()
	defer seg.RUnlock()
	value, ok := sm.m[len(key)%numSegments][key]
	return value, ok
}

func main() {
	sm := NewSafeMap()

	sm.Set("key1", 1)
	value, ok := sm.Get("key1")
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
