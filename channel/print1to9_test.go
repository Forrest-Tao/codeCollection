package main

import (
	"sync"
	"testing"
)

func TestPrintEven(t *testing.T) {
	var (
		wg   sync.WaitGroup
		even = make(chan struct{}, 1)
		old  = make(chan struct{}, 1)
	)

	wg.Add(2)

	PrintEven(&wg, even, old)
	PrintOld(&wg, even, old)
	even <- struct{}{}
	wg.Wait()
}
