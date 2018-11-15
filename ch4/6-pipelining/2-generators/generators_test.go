package generators

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRepeat(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}

func TestRepeatFn(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	rand := func() interface{} { return rand.Int() }
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Printf("%v ", num)
	}
}

func TestToString(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}
