package dep_test

import (
	"log"
	"runtime"
	"sync"
	"testing"
)

func runMultiSender() {
	log.SetFlags(0)

	const Max = 1000
	const NumSenders = 10
	const NumReceivers = 1
	wgMain := sync.WaitGroup{}
	wgMain.Add(NumSenders + NumReceivers)

	dataCh := make(chan int)
	stopCh := make(chan struct{})

	sumSender := 0
	sumSendItems := make([]int, NumSenders)
	wgSenders := sync.WaitGroup{}
	wgSenders.Add(NumSenders)
	for i := NumSenders - 1; i >= 0; i-- {
		go func(i int) {
			defer wgMain.Done()

			for j := Max; ; j-- {
				value := 0
				if j == 0 {
					wgSenders.Done()
					wgSenders.Wait()
					sumSender += sumSendItems[i]
					log.Println(i, sumSendItems[i], sumSender)
				} else if j > 0 {
					value = i*Max + j
					sumSendItems[i] += value
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(i)
	}

	sumReceiver := 0
	go func() {
		defer wgMain.Done()

		for value := range dataCh {
			if value == 0 {
				// 唯一的接收者可以安全地关闭stopCh通道
				close(stopCh)
				return
			}

			sumReceiver += value
		}
	}()
	wgMain.Wait()
	log.Println(sumSender, sumSendItems, sumReceiver)
}

func TestRunMultiSender(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	runMultiSender()
	t.Log(runtime.NumGoroutine())
}
