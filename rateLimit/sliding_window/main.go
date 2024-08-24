package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mutex    sync.Mutex
	interval time.Duration
	segments int
	limit    int
	windows  []int
	start    time.Time
}

func NewSlidingWindowLimiter(interval time.Duration, segments, limit int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		interval: interval,
		segments: segments,
		limit:    limit,
		windows:  make([]int, segments),
		start:    time.Now(),
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(s.start)
	segmentDuration := s.interval / time.Duration(s.segments)

	steps := int(elapsed / time.Duration(s.segments))
	if steps > 0 {
		s.slideWindow(steps)
		s.start = s.start.Add(time.Duration(steps) * segmentDuration)
	}
	currentCount := 0
	for _, count := range s.windows {
		currentCount += count
	}
	if currentCount > s.limit {
		return false
	}
	s.windows[len(s.windows)-1]++
	return true
}

func (s *SlidingWindowLimiter) slideWindow(steps int) {
	if steps > s.segments {
		for i := 0; i < s.segments; i++ {
			s.windows[i] = 0
		}
	} else {
		copy(s.windows, s.windows[steps:])
		for i := len(s.windows) - steps; i < len(s.windows); i++ {
			s.windows[i] = 0
		}
	}
}

func main() {
	limiter := NewSlidingWindowLimiter(time.Minute/10, 12, 80)

	for i := 0; i < 105; i++ {
		if limiter.Allow() {
			//fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(5 * time.Second / 1000) // 模拟请求间隔
	}
}
