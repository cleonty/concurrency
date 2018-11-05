package main

import "fmt"

func main() {
	incStream := make(chan int)
	go func() {
		defer close(incStream)
		for i := 1; i <= 5; i++ {
			incStream <- i
		}
	}()

	for integer := range incStream {
		fmt.Printf("%v ", integer)
	}
}
