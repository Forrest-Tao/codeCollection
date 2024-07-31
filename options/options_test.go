package options

import (
	"fmt"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	server := NewServer(
		WithAddress("192.168.1.1"),
		WithPort(9090),
		WithReadTimeout(10*time.Second),
	)

	fmt.Printf("Server: %+v\n", server)
}
