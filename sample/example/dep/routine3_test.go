package dep_test

import (
	"log"
	"runtime"
	"sync"
	"testing"
)

func runMultiReceiver() {
	log.SetFlags(0)

	dataCh := make(chan int)

	const Max = 10000
	sumSender := 0
	go func() {
		for value := Max; ; value-- {
			if value == 0 {
				// 唯一的发送者可以安全地关闭dataCh通道
				close(dataCh)
				return
			}

			dataCh <- value
			sumSender += value
		}
	}()

	const NumReceivers = 10
	sumReceiver := 0
	sumReceiveItems := make([]int, NumReceivers)
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	for i := 0; i < NumReceivers; i++ {
		go func(i int) {
			defer wgReceivers.Done()

			// 多个接收者互相竞争地接收数据，直到通道队列已空且被关闭
			for value := range dataCh {
				sumReceiveItems[i] += value
			}

			sumReceiver += sumReceiveItems[i]
		}(i)
	}
	wgReceivers.Wait()
	log.Println(sumSender, sumReceiver, sumReceiveItems)
}

func TestRunMultiReceiver(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	runMultiReceiver()
	t.Log(runtime.NumGoroutine())
}
