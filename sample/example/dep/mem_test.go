// generate_test.go
package dep_test

import (
	"math/rand"
	"testing"
	"time"
)

func generate(n int, hasCap bool) []int {
	rand.Seed(time.Now().UnixNano())
	var nums []int
	if hasCap {
		nums = make([]int, 0, n)
	} else {
		nums = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BmGenerate(n int, hasCap bool, b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(n, hasCap)
	}
}

func BenchmarkGenerateNoCap1(b *testing.B)        { BmGenerate(1, false, b) }
func BenchmarkGenerateNoCap10(b *testing.B)       { BmGenerate(10, false, b) }
func BenchmarkGenerateNoCap100(b *testing.B)      { BmGenerate(100, false, b) }
func BenchmarkGenerateNoCap1000(b *testing.B)     { BmGenerate(1000, false, b) }
func BenchmarkGenerateNoCap10000(b *testing.B)    { BmGenerate(10000, false, b) }
func BenchmarkGenerateNoCap100000(b *testing.B)   { BmGenerate(100000, false, b) }
func BenchmarkGenerateNoCap1000000(b *testing.B)  { BmGenerate(1000000, false, b) }
func BenchmarkGenerateHasCap1000000(b *testing.B) { BmGenerate(1000000, true, b) }
