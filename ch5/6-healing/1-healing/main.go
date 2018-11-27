package main

import (
	"log"
	"os"
	"time"
)

func doWork(done <-chan interface{}, _ time.Duration) <-chan interface{} {
	log.Println("ward: Hello, I'm irresponsible!")
	go func() {
		<-done
		log.Println("ward: I am halting.")
	}()
	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)
	doWorkWithSteward := newSteward(4*time.Second, doWork)
	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})
	for range doWorkWithSteward(done, 4*time.Second) {
	}
	log.Println("done")
}
