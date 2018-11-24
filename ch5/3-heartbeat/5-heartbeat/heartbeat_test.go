package heartbeat

import (
	"testing"
	"time"
)

func TestDoWork_GenerateAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4}
	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2, intSlice...)
	<-heartbeat
	i := 0
	for {
		select {
		case r, ok := <-results:
			if !ok {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
			}
			i++
		case <-heartbeat:
		case <-time.After(timeout):
			t.Fatal("test timed out")
		}
	}
}
