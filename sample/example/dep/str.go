package dep

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ConcatRandomString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += RandomString(n)
	}
	return s
}

func ConcatRandomStringByBuilder(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(RandomString(n))
	}
	return s.String()
}
