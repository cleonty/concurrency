package main

import (
	"fmt"
)

func fib(n int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		if n <= 2 {
			result <- 1
			return
		}
		result <- <-fib(n-1) + <-fib(n-2)
	}()
	return result
}

func main() {
	const n = 4
	fmt.Printf("fib(%d) = %d\n", n, <-fib(n))
}
