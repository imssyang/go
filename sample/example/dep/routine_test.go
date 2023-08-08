package dep_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doBad(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func doGood(waitDone, done chan bool) {
	time.Sleep(time.Second)
	select {
	case waitDone <- true:
	default:
		return
	}
	time.Sleep(time.Second)
	done <- true
}

func runBad(size int) error {
	// BAD: unbufferd chan will block forever after timeout!
	done := make(chan bool, size)
	go doBad(done)

	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func runGood() error {
	waitDone := make(chan bool)
	done := make(chan bool)
	go doGood(waitDone, done)

	select {
	case <-waitDone:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func testRun(t *testing.T, size int, args ...any) {
	t.Helper()
	t.Log("S>", runtime.NumGoroutine())
	for i := 0; i < size; i++ {
		switch f := args[0].(type) {
		case func() error:
			f()
		case func(int) error:
			f(args[1].(int))
		default:
			fmt.Printf("%T", size)
		}
	}
	time.Sleep(time.Second * 3)
	t.Log("E>", runtime.NumGoroutine())
}

func TestTimeout1(t *testing.T) { testRun(t, 1000, runBad, 0) }
func TestTimeout2(t *testing.T) { testRun(t, 1000, runBad, 1) }
func TestTimeout3(t *testing.T) { testRun(t, 1000, runGood) }
