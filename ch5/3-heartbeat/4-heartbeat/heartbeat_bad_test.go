package main

import (
	"testing"
	"time"
)

func TestDoWorkBad_GenerateAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	_, results := doWork(done, intSlice...)
	for i, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("test timed out")
		}
	}
}
