package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 判断channel是否关闭 并发不安全

type T int

func IsClosed(t <-chan T) bool {
	select {
	case <-t:
		return true
	default:
	}
	return false
}

// 一个接受者 多个发送者

func ReceiverCloseChan() {
	rand.Seed(time.Now().Unix())

	const (
		Max     = 10000
		Senders = 100
	)

	wgReceiver := sync.WaitGroup{}
	wgReceiver.Add(1)
	dataChan := make(chan int)
	closeChan := make(chan struct{})

	for i := 0; i < Senders; i++ {
		go func() {
			for {
				select {
				case <-closeChan:
					return
				default:
				}

				select {
				case <-closeChan:
					return
				case dataChan <- rand.Intn(Max):
				}
			}
		}()
	}

	go func() {
		defer func() {
			wgReceiver.Done()
		}()

		for v := range dataChan {
			if v == Max-1 {
				close(closeChan)
				return
			}
			fmt.Println(v)
		}
	}()

	wgReceiver.Wait()
}
