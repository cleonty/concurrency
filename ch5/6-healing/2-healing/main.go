package main

import (
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
					log.Printf("next intList value is %d\n", intValue)
					if intValue < 0 {
						log.Printf("negative value: %v\n", intValue)
						return
					}
					for {
						select {
						case <-pulse:
							log.Printf("doWork: about to send heartbeat")
							select {
							case heartbeat <- struct{}{}:
								log.Printf("doWork: sent heartbeat")
							default:
								log.Printf("doWork: not sent heartbeat")
							}
						case intStream <- intValue:
							log.Printf("doWork: sent intValue %d\n", intValue)
							continue valueLoop
						case <-done:
							log.Printf("doWork: done closed; exiting")
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
		log.Printf("main received: %v\n", intVal)
	}
	log.Println("done")
}
