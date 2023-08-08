package dep

import (
	"fmt"
	"time"
)

func DoBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func Timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
