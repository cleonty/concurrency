package main

import (
	"log"
	"time"
)

func doWork(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)

	go func() {
		defer close(heartbeat)
		defer close(results)

		pulse := time.Tick(pulseInterval)
		workGen := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}

		sendResult := func(r time.Time) {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case results <- r:
				return
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}
	}()
	return heartbeat, results
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		log.Println("about to close done channel")
		close(done)
		log.Println("done channel closed")
	})
	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				log.Println("exit because heatbeat channel was closed")
				return
			}
			log.Println("Pulse")
		case r, ok := <-results:
			if !ok {
				log.Println("exit because results channel was closed")
				return
			}
			log.Printf("results %v\n", r.Second())
		case <-time.After(timeout):
			log.Println("exit because worker goroutine is not healthy!")
			return
		}
	}
}
