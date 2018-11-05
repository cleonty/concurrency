package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	incStream := make(chan int, 4)
	go func() {
		defer close(incStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			incStream <- i
		}
	}()
	for integer := range incStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
