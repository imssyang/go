package dep_test

import (
	"log"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func runMultiSenderAndReceiver() {
	log.SetFlags(0)

	const Max = 1000
	const NumSenders = 10
	const NumReceivers = 10
	const NumAgents = 1
	wgMain := sync.WaitGroup{}
	wgMain.Add(NumSenders + NumReceivers + NumAgents)

	dataCh := make(chan int)
	stopCh := make(chan struct{})
	agentCh := make(chan string, 1)

	var stopper string
	go func() {
		defer wgMain.Done()

		stopper = <-agentCh

		// 唯一的调解者可以安全地关闭stopCh通道
		close(stopCh)
	}()

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

					// 向中间调解者发送信号（非阻塞）
					select {
					case agentCh <- "Sender#" + strconv.Itoa(i):
					default:
					}
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
	sumReceiveItems := make([]int, NumReceivers)
	for i := 0; i < NumReceivers; i++ {
		go func(i int) {
			defer wgMain.Done()

			for {
				select {
				case <-stopCh:
					sumReceiver += sumReceiveItems[i]
					return
				case value := <-dataCh:
					if value == 0 {
						// 向中间调解者发送信号（非阻塞）
						select {
						case agentCh <- "Receiver#" + strconv.Itoa(i):
						default:
						}
					} else {
						sumReceiveItems[i] += value
					}
				}
			}
		}(i)
	}
	wgMain.Wait()
	log.Println(stopper, sumSender, sumSendItems, sumReceiver, sumReceiveItems)
}

func TestRunMultiSenderAndReceiver(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	runMultiSenderAndReceiver()
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
