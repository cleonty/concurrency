package main

import "testing"

func TestWrite(t *testing.T) {
	var dataStream chan interface{}
	dataStream <- struct{}{}
}

func TestRead(t *testing.T) {
	var dataStream chan interface{}
	<-dataStream
}
func TestClose(t *testing.T) {
	var dataStream chan interface{}
	close(dataStream)
}
