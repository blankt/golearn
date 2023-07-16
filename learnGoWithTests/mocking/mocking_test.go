package mocking

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	writer := &bytes.Buffer{}
	sleeper := &ConfigurableSleeper{
		duration: time.Second * 1,
		sleep:    time.Sleep,
	}
	CountDown(writer, sleeper)

	result := writer.String()
	point := `3
2
1
Go`
	if result != point {
		t.Fatal(fmt.Sprintf("want %s get %s", point, result))
	}

	t.Run("test count down operations", func(t *testing.T) {
		spyOperations := &SpyCountDownOperations{}
		CountDown(spyOperations, spyOperations)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spyOperations.call) {
			t.Fatal(fmt.Sprintf("want %v but get %v", want, spyOperations.call))
		}
	})
}
