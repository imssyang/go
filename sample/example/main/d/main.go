package main

import "fmt"

func Increase() func() int {
	n := 0 // n 会一直存在，直到 in 被销毁（逃逸到heap）
	return func() int {
		n++
		return n
	}
}

func main() {
	in := Increase()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2
}

/*
$ go build -gcflags=-m ./main/d
# example.com/main/d
...
main/d/main.go:6:2: moved to heap: n
main/d/main.go:7:9: func literal escapes to heap
main/d/main.go:14:16: func literal does not escape
main/d/main.go:15:13: ... argument does not escape
main/d/main.go:15:16: ~R0 escapes to heap
main/d/main.go:16:13: ... argument does not escape
main/d/main.go:16:16: ~R0 escapes to heap
*/
