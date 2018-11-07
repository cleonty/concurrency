package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer log.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	log.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		log.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}
