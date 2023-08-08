package dep_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doThing(data chan int) {
	for {
		select {
		// ok=true表示数据是有效载荷，false在chan关闭后产生
		// chan已经关闭后，<-ch 将不会阻塞
		case v, ok := <-data:
			if !ok {
				fmt.Println("data has been closed")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", v)
		}
	}
}

func runExitByClose() {
	data := make(chan int, 10)
	go doThing(data)
	for i := 0; i < 1000; i++ {
		data <- i
	}
	close(data)
}

func TestExitByClose(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	runExitByClose()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
