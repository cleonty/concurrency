package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func doWorkFn(done <-chan interface{}, intList ...int) (startGoroutineFn, <-chan interface{}) {
	intChanStream := make(chan (<-chan interface{}))
	intStream := bridge(done, intChanStream)
	doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		intStream := make(chan interface{})
		heartbeat := make(chan interface{})
		go func() {
			defer close(intStream)
			select {
			case intChanStream <- intStream:
			case <-done:
				return
			}
			pulse := time.Tick(pulseInterval)
			for {
			valueLoop:
				for _, intValue := range intList {
					if intValue < 0 {
						log.Printf("negative value: %v\n", intValue)
						return
					}
					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case intStream <- intValue:
							continue valueLoop
						case <-done:
							return
						}
					}
				}
			}

		}()
		return heartbeat
	}
	return doWork, intStream
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)
	done := make(chan interface{})
	defer close(done)
	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)
	for intVal := range take(done, intStream, 6) {
		fmt.Printf("Received: %v\n", intVal)
	}
	log.Println("done")
}
