package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v) %v\n", ok, salutation)
}
