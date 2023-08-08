package dep_test

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [4096]byte
}

var j, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := &Student{}
		json.Unmarshal(j, s)
	}
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func BenchmarkUnmarshalByPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := pool.Get().(*Student)
		json.Unmarshal(j, s)
		pool.Put(s)
	}
}
