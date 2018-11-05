package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello channels!"
	}()
	fmt.Println(<-stringStream)
}
