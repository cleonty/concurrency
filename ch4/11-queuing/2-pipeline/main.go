package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	zeros := take(done, 3, repeat(done, 0))
	short := sleep(done, 1*time.Second, zeros)
	buffer := buffer(done, 2, short) // Buffers sends from short by 2
	long := sleep(done, 4*time.Second, buffer)
	pipeline := long
	start := time.Now()
	for range pipeline {

	}
	fmt.Printf("Elapsed: %s\n", time.Since(start))
}
