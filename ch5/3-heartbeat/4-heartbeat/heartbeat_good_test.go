package main

import (
	"testing"
)

func TestDoWorkGood_GenerateAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	heartbeat, results := doWork(done, intSlice...)

	<-heartbeat

	i := 0

	for r := range results {
		if r != intSlice[i] {
			t.Errorf("index %v: expected %v, but received %v,", i, intSlice[i], r)
		}
		i++
	}
}
