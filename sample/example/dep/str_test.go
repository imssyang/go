package dep_test

import (
	"fmt"
	"strings"
	"testing"

	"example.com/dep"
)

func TestConcatByBuilder(t *testing.T) {
	var str = dep.RandomString(10)
	var builder strings.Builder
	cap := 0
	for i := 0; i < 10000; i++ {
		if builder.Cap() != cap {
			fmt.Print(builder.Cap(), " ")
			cap = builder.Cap()
		}
		builder.WriteString(str)
	}
}
