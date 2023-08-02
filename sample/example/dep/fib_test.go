package dep_test

import (
	"testing"
	"time"

	"example.com/dep"
)

func BenchmarkFib(b *testing.B) {
	time.Sleep(time.Second * 3)
	b.ResetTimer() // 消除上面sleep的影响
	for n := 0; n < b.N; n++ {
		dep.Fib(30) // run fib(30) b.N times
	}
}
