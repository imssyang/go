package dep_test

import (
	"bytes"
	"sync"
	"testing"
)

var data = make([]byte, 10000)

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func BenchmarkBufferByPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufPool.Put(buf)
	}
}
