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
		for {
			if m1.TryLock() {
				break
			}
			time.Sleep(time.Millisecond)
		}
		defer m1.Unlock()
		time.Sleep(time.Second)
		for {
			if m2.TryLock() {
				break
			}
			time.Sleep(time.Millisecond)
		}
		defer m2.Unlock()
		fmt.Println("goroutine 1")
	}()
	go func() {
		for {
			if m2.TryLock() {
				break
			}
			time.Sleep(time.Millisecond)
		}
		defer m2.Unlock()
		time.Sleep(time.Second)
		for {
			if m1.TryLock() {
				break
			}
			time.Sleep(time.Millisecond)
		}
		defer m1.Unlock()
		fmt.Println("goroutine 2")
	}()
	time.Sleep(5 * time.Second)
}
