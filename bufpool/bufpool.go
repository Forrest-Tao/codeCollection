package bufpool

import (
	"bytes"
	"sync"
)

var pools = createPools([]int{128, 512, 1024, 2048, 4096, 8192, 16 * 1024, 32 * 1024, 64 * 1024, 65 * 1024})

var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	},
}

func createPools(sizes []int) []struct {
	size int
	pool sync.Pool
} {
	p := make([]struct {
		size int
		pool sync.Pool
	}, len(sizes))

	for i, size := range sizes {
		s := size
		p[i] = struct {
			size int
			pool sync.Pool
		}{
			size: s,
			pool: sync.Pool{
				New: func() any {
					b := make([]byte, s)
					return &b
				},
			},
		}
	}
	return p
}

// Get returns a buffer of specified size.
func Get(size int) *[]byte {
	for i := range pools {
		if size <= pools[i].size {
			b := pools[i].pool.Get().(*[]byte)
			*b = (*b)[:size]
			return b
		}
	}
	b := make([]byte, size)
	return &b
}

// Put returns a buffer to the pool.
func Put(b *[]byte) {
	for i := range pools {
		if cap(*b) == pools[i].size {
			pools[i].pool.Put(b)
			return
		}
	}
}

// GetBuff returns a buffer from the pool.
func GetBuff() *bytes.Buffer {
	buffer := bufPool.Get().(*bytes.Buffer)
	buffer.Reset()
	return buffer
}

// PutBuff returns a buffer to the pool.
func PutBuff(buf *bytes.Buffer) {
	buf.Reset()
	bufPool.Put(buf)
}
