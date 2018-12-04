package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	// Every second, log how many goroutines are currently running.
	go func() {
		goroutines := pprof.Lookup("goroutine")
		threads := pprof.Lookup("threadcreate")
		for range time.Tick(1 * time.Second) {
			log.Printf("goroutine count: %d\n", goroutines.Count())
			log.Printf("thread count: %d\n", threads.Count())
		}
	}()

	var blockForever chan struct{}
	for i := 0; i < 10; i++ {
		go func() { <-blockForever }()
		time.Sleep(500 * time.Millisecond)
	}
}
