package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

// 以下代码会导致死锁
func broadcast() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			rand.Seed(time.Now().Unix())
			nums := rand.Intn(100)
			time.Sleep(time.Millisecond * time.Duration(nums))
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoyling dialog box")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
