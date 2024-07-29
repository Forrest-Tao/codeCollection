package pool

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name  string
	Age   int
	bytes [2048]byte
}

var stuPool = sync.Pool{New: func() interface{} {
	return new(Student)
}}

var buff, _ = json.Marshal(Student{
	Name:  "xx",
	Age:   123,
	bytes: [2048]byte{},
})

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := stuPool.Get().(*Student)
		json.Unmarshal(buff, stu)
		stuPool.Put(stu)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := &Student{}
		json.Unmarshal(buff, stu)
	}
}
