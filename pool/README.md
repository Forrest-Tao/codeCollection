
```bash
❯ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: poolDemo
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkBufferWithPool-12              14562392                82.86 ns/op            0 B/op          0 allocs/op
BenchmarkBuffer-12                        992318              1156 ns/op           10240 B/op          1 allocs/op
BenchmarkUnmarshalWithPool-12            2196350               545.3 ns/op           224 B/op          5 allocs/op
BenchmarkUnmarshal-12                    1447239               833.1 ns/op          2528 B/op          6 allocs/op
PASS
ok      poolDemo        6.254s
```

## 案例
**源码位置/src/fmt/print.go**

```go
var ppFree = sync.Pool{
	New: func() any { return new(pp) },
}

// newPrinter allocates a new pp struct or grabs a cached one.
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.wrapErrs = false
	p.fmt.init(&p.buf)
	return p
}

// free saves used pp structs in ppFree; avoids an allocation per invocation.
func (p *pp) free() {
	// Proper usage of a sync.Pool requires each entry to have approximately
	// the same memory cost. To obtain this property when the stored type
	// contains a variably-sized buffer, we add a hard limit on the maximum
	// buffer to place back in the pool. If the buffer is larger than the
	// limit, we drop the buffer and recycle just the printer.
	//
	// See https://golang.org/issue/23199.
	if cap(p.buf) > 64*1024 {
		p.buf = nil
	} else {
		p.buf = p.buf[:0]
	}
	if cap(p.wrappedErrs) > 8 {
		p.wrappedErrs = nil
	}

	p.arg = nil
	p.value = reflect.Value{}
	p.wrappedErrs = p.wrappedErrs[:0]
	ppFree.Put(p)
}

```