package main

import (
	"fmt"
)

func main() {
	myChan := take(nil, repeat(nil, 1, 2, 3, 4, 5), 50)
	done := make(chan interface{})
	defer close(done)
	for val := range orDone(done, myChan) {
		fmt.Printf("%v\n", val)
	}
}
