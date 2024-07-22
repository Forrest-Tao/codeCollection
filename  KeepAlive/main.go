package main

import "time"

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
