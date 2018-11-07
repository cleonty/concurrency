package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer log.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					log.Printf("Do something interesting with %q\n", s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		log.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated
	log.Println("Done.")
}
