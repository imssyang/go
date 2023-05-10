package main

import "sync"

var l sync.Mutex
var a string

func f() {
	a = "Hi"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
