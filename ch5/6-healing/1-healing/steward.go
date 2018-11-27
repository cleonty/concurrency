package main

import (
	"log"
	"time"
)

type startGoroutineFn func(cone <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})

func newSteward(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
	return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
			}
			startWard()
			pulse := time.Tick(pulseInterval)
		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)
				for {
					select {
					case <-pulse:
						log.Printf("pulse; about to send a heartbeat")
						select {
						case heartbeat <- struct{}{}:
							log.Printf("heartbeat sent")
						default:
							log.Printf("heartbeat not sent")
						}
					case <-wardHeartbeat:
						log.Printf("got ward's heartbeat, continue to monitor")
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward is unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						log.Printf("done is closed; exiting")
						return
					}
				}
			}
		}()
		return heartbeat
	}
}
