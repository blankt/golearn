package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerBits   = 10
	sequenceBits = 12
	workerMax    = -1 ^ (-1 << workerBits)
	sequenceMask = -1 ^ (-1 << sequenceBits)
	timeShift    = workerBits + sequenceBits
	workerShift  = sequenceBits
	epoch        = 1609459200000
)

type Snowflake struct {
	mu       sync.Mutex
	lastTime int64
	workerID int64
	sequence int64
}

func NewSnowflake(workerID int64) *Snowflake {
	if workerID < 0 || workerID > workerMax {
		panic(fmt.Sprintf("worker ID must be between 0 and %d", workerMax))
	}
	return &Snowflake{
		lastTime: time.Now().UnixNano() / 1e6,
		workerID: workerID,
		sequence: 0,
	}
}

func (sf *Snowflake) NextID() int64 {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	currentTime := time.Now().UnixNano() / 1e6

	if currentTime < sf.lastTime {
		panic(fmt.Sprintf("clock moved backwards, refusing to generate ID for %d milliseconds", sf.lastTime-currentTime))
	}

	if currentTime == sf.lastTime {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			for currentTime <= sf.lastTime {
				currentTime = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTime = currentTime

	id := (currentTime-epoch)<<timeShift | (sf.workerID << workerShift) | sf.sequence
	return id
}

func main() {
	sf := NewSnowflake(1) // 假设工作节点ID为1

	for i := 0; i < 10; i++ {
		id := sf.NextID()
		fmt.Println(id)
		//time.Sleep(time.Millisecond)
	}
	fmt.Println("generate down")
}
