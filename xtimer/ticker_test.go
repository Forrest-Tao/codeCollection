package xtimer

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	now := time.Now()
	_timer := time.NewTimer(0)
	<-_timer.C
	fmt.Println("Timer triggered", "after", time.Since(now))
}
