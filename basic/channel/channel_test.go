package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestIsClosed(t *testing.T) {
	closeChan := make(chan T)
	ok := IsClosed(closeChan)
	if ok {
		t.Fatal("is closed func error")
	}
	close(closeChan)
	ok = IsClosed(closeChan)
	if !ok {
		t.Fatal("is closed func error")
	}
}

func TestSafeClose(t *testing.T) {
	closeChan := make(chan T)
	close(closeChan)
	closed := SafeClosed(closeChan)
	if !closed {
		t.Fatal("SafeClosed func error")
	}
}

// 已关闭的channel仍能取到数据 为channel类型的默认0值 可以用此来做同步信号
func TestChannel(t *testing.T) {
	testChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 3)
		close(testChan)
		wg.Done()
	}()

	wg.Wait()
	for i := 0; i < 10; i++ {
		select {
		case x := <-testChan:
			fmt.Printf("从已关闭的chan中取到数据:%v \n", x)
		}
	}
}

// 遍历channel取数据 当channel关闭后也会把其中的数据取完再结束遍历
func TestCloseChannel(t *testing.T) {
	testChan := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			testChan <- i
		}
		close(testChan)
		wg.Done()
	}()

	wg.Wait()
	for v := range testChan {
		fmt.Println(v)
	}
}

func TestOneReceiver(t *testing.T) {
	ReceiverCloseChan()
}

func TestMReceiverOneSender(t *testing.T) {
	MReceiverOneSender()
}

func TestMReceiverMSender(t *testing.T) {
	MReceiverMSender()
}

// 阻塞式的通道 后面发送的数据会覆盖前一个
func TestSend(t *testing.T) {
	dataChan := make(chan int)
	for i := 0; i < 2; i++ {
		go func(index int) {
			if index == 1 {
				time.Sleep(time.Second * 1)
			}
			dataChan <- index
		}(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(<-dataChan)
}

func TestThirdGoroutineClose(t *testing.T) {
	ThirdGoroutineClose()
}

func TestReceiverCloseChan2(t *testing.T) {
	ReceiverCloseChan2()
}
