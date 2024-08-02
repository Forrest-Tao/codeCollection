package singleton

import "testing"

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		GetInstance()
	}
	println(cnt == 1)
}
