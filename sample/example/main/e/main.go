package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("Goroutine 1:", value)
		}
	}()

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("Goroutine 2:", value)
		}
	}()

	wg.Wait()
}
