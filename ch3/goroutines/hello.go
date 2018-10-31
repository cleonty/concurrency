package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello() {
	defer wg.Done()
	fmt.Printf("hello")
}

func main() {
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
