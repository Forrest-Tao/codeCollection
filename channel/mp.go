package main

import (
	"context"
	"sync"
	"time"
)

/*
1.实现一个并发的map
2.实现O（1）插入和查询
3.查询存在，则直接返回；查询不存在，阻塞知道获取到val，返回，若指定时间后还未获取到，直接返回超时错误
4.不能出现panic
*/
type MyChan struct {
	ch chan struct{}
	sync.Once
}

func NewMyChan() MyChan {
	return MyChan{
		ch: make(chan struct{}),
	}
}

func (m *MyChan) Close() {
	m.Once.Do(
		func() {
			close(m.ch)
		})
}

type MyConcurrentMap struct {
	mp map[int]int
	sync.Mutex
	keyToCh map[int]MyChan
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp:      make(map[int]int),
		keyToCh: make(map[int]MyChan),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v
	if ch, ok := m.keyToCh[k]; ok {
		ch.Close()
		delete(m.keyToCh, k)
	}
}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	if v, ok := m.mp[k]; ok {
		m.Unlock()
		return v, nil
	}
	ch, ok := m.keyToCh[k]
	if !ok {
		ch = NewMyChan()
		m.keyToCh[k] = ch
	}
	ctx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()

	select {
	case <-ctx.Done():
		delete(m.keyToCh, k)
		m.Unlock()
		return -1, ctx.Err()
	case <-ch.ch:
		m.Unlock()
		return m.mp[k], nil
	}
}
