package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chanOwner := func(wg *sync.WaitGroup, done <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			defer wg.Done()
			for _, s := range []string{"a", "b", "c"} {
				select {
				case <-time.After(1 * time.Second):
				case <-done:
					fmt.Println("done")
					return
				case stringStream <- s:
				}
			}
		}()
		return stringStream
	}
	consumer := func(wg *sync.WaitGroup, stringStream <-chan string) {
		defer wg.Done()
		for s := range stringStream {
			fmt.Println(s)
		}
	}
	var wg sync.WaitGroup
	wg.Add(2)
	done := make(chan interface{})
	stringStream := chanOwner(&wg, done)
	go consumer(&wg, stringStream)
	wg.Wait()
}
