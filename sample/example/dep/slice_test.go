package dep_test

import (
	"math/rand"
	"runtime"
	"testing"
	"time"

	"example.com/dep"
)

func init1MB() []int {
	rand.Seed(time.Now().UnixNano())
	n := 128 * 1024 //1M in x64 machine
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func printMem(t *testing.T) {
	t.Helper()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t.Logf("%.2f MB", float64(m.Alloc)/1024./1024.)
}

func testF(t *testing.T, f func([]int) []int) {
	t.Helper()
	r := make([][]int, 0)
	for i := 0; i < 100; i++ {
		origin := init1MB()
		r = append(r, f(origin))
		// runtime.GC() // (maybe)
	}
	printMem(t)
	_ = r
}

func TestLastNumsBySlice(t *testing.T) { testF(t, dep.LastNumsBySlice) }
func TestLastNumsByCopy(t *testing.T)  { testF(t, dep.LastNumsByCopy) }

/*
❯ go test -run="TestLast" . -v                                                                                            master ✖ ◼
=== RUN   TestLastNumsBySlice slice_test.go:40: 100.12 MB
=== RUN   TestLastNumsByCopy slice_test.go:41: 2.13 MB
*/
