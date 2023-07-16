package mocking

import (
	"fmt"
	"io"
	"time"
)

//TDD so important,must be learned more.

const (
	BeginNum  = 3
	FinalWord = "Go"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := BeginNum; i > 0; i-- {
		_, _ = fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprint(writer, FinalWord)
}

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountDownOperations struct {
	call []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.call = append(s.call, sleep)
}

func (s *SpyCountDownOperations) Write(data []byte) (n int, err error) {
	s.call = append(s.call, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
