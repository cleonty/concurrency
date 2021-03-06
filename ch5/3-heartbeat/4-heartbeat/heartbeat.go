package main

import (
	"fmt"
	"log"
	"time"
)

func doWork(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeatStream := make(chan interface{}, 1)
	workStream := make(chan int)

	go func() {
		defer close(heartbeatStream)
		defer close(workStream)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			fmt.Println("begin unit of work")
			select {
			case heartbeatStream <- struct{}{}:
				fmt.Println("heartbeat sent")
			default:
				fmt.Println("heartbeat didn't sent because previous heartbeat hasn't read yet")
			}
			select {
			case <-done:
				log.Printf("worker goroutine extits because done channel has closed")
				return
			case workStream <- n:
				fmt.Println("sent new unit of work")
			}
		}
	}()
	return heartbeatStream, workStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done, 1, 2, 3, 4)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("Pulse")
			} else {
				fmt.Println("exit because no pulse, heartbeat channel has been closed")
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				fmt.Println("exit because results channel has been closed")
				return
			}
		}
	}
}
