package channel

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

func SafeClosed(t chan T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
			return
		}
	}()
	close(t)
	return
}

// MReceiverOneSender 多个接收者 一个发送者
func MReceiverOneSender() {
	rand.Seed(time.Now().Unix())
	const (
		Max         = 100000
		NumReceiver = 100
	)
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceiver)
	dataCh := make(chan int)

	// 发送者
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	//接收者
	for i := 0; i < NumReceiver; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

// ReceiverCloseChan 一个接受者 多个发送者
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

// MReceiverMSender 多个发送者 多个接受者
func MReceiverMSender() {
	rand.Seed(time.Now().Unix())
	const (
		Max       = 10000
		Senders   = 1000
		Receivers = 10
	)
	var (
		stopBy string
	)

	wgReceiver := sync.WaitGroup{}
	wgReceiver.Add(Receivers)

	dataChan := make(chan int)
	//toStop 容量设置为1是为了确保在关闭stopChan协程关闭stopChan之前,从toSop中取到的值是第一个发送信号的协程
	toStop := make(chan string, 1)
	stopChan := make(chan struct{})

	//守护者协程 用于通知接受者、发送者通道已关闭
	go func() {
		stopBy = <-toStop
		close(stopChan)
	}()

	for i := 0; i < Senders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStop <- "sender " + id:
					default:
					}
					return
				}

				select {
				case <-stopChan:
					return
				default:
				}

				select {
				case <-stopChan:
					return
				case dataChan <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	for i := 0; i < Receivers; i++ {
		go func(id string) {
			defer wgReceiver.Done()
			for {
				value := rand.Intn(Max)
				if value == Max-1 {
					select {
					case toStop <- "receiver " + id:
					default:
					}
					return
				}

				select {
				case <-stopChan:
					return
				default:
				}

				select {
				case <-stopChan:
					return
				case data := <-dataChan:
					fmt.Println(data)
				}
			}
		}(strconv.Itoa(i))
	}

	wgReceiver.Wait()
	log.Printf("stop by %s", stopBy)
}

func ThirdGoroutineClose() {
	rand.Seed(time.Now().Unix())

	const (
		Max            = 10000
		Receivers      = 10
		ThirdGoroutine = 5
	)

	dataChan := make(chan int)
	closing := make(chan struct{})
	closed := make(chan struct{})
	wgReceiver := sync.WaitGroup{}
	wgReceiver.Add(Receivers)

	stopped := func() {
		select {
		case closing <- struct{}{}:
			<-closed
		case <-closed:
		}
	}

	for i := 0; i < ThirdGoroutine; i++ {
		go func() {
			r := rand.Intn(3) + 1
			time.Sleep(time.Duration(r) * time.Second)
			stopped()
		}()
	}

	go func() {
		defer func() {
			close(dataChan)
			close(closed)
		}()
		for {
			value := rand.Intn(Max)

			select {
			case <-closing:
				return
			default:
			}
			dataChan <- value
		}
	}()

	for i := 0; i < Receivers; i++ {
		go func() {
			defer wgReceiver.Done()
			for data := range dataChan {
				fmt.Println(data)
			}
		}()
	}
	wgReceiver.Wait()
}

func ReceiverCloseChan2() {
	rand.Seed(time.Now().UnixNano())

	const Max = 1000000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int)
	middleCh := make(chan int)
	closing := make(chan struct{})
	closed := make(chan struct{})

	go func() {
		for {
			select {
			case <-closing:
				close(closed)
				close(dataCh)
				return
			case v := <-middleCh:
				select {
				case <-closing:
					close(closed)
					close(dataCh)
					return
				case dataCh <- v:
				}
			}
		}
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case closing <- struct{}{}:
					case <-closed:
						return
					}
					return
				}

				select {
				case <-closed:
					return
				default:
				}

				select {
				case <-closed:
					return
				case middleCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for range [NumReceivers]struct{}{} {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
