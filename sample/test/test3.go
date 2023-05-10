package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1 sync.Mutex
	m2 sync.Mutex
)

func main() {
	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(time.Second)
		m2.Lock()
		defer m2.Unlock()
		fmt.Println("goroutine 1")
	}()
	go func() {
		m2.Lock()
		defer m2.Unlock()
		time.Sleep(time.Second)
		m1.Lock()
		defer m1.Unlock()
		fmt.Println("goroutine 2")
	}()
	time.Sleep(2 * time.Second)
}
