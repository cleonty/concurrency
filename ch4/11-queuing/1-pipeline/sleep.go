package main

import (
	"time"
)

func sleep(done <-chan interface{}, d time.Duration, valueStream <-chan interface{}) <-chan interface{} {
	delayedStream := make(chan interface{})
	go func() {
		defer close(delayedStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case <-time.After(d):
				select {
				case <-done:
					return
				case delayedStream <- v:
				}
			}
		}
	}()
	return delayedStream
}
