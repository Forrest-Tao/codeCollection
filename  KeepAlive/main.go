package main

import (
	"fmt"
	"time"
)

func keepAlive() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for _ = range ticker.C {

		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:

			}
		}
	}()
}

func main() {
	cnt := 1000
	var sum time.Duration
	for i := 0; i < 2; i++ {
		now := time.Now()
		time.Sleep(time.Second)
		sum += time.Since(now)
	}
	fmt.Printf("sum: %f\n", sum.Seconds())
	fmt.Printf("qps: %d", cnt/int(sum.Seconds()))
}
