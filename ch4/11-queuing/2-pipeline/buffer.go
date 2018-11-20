package main

func buffer(done <-chan interface{}, num int, valueStream <-chan interface{}) <-chan interface{} {
	bufferedStream := make(chan interface{}, num)
	go func() {
		defer close(bufferedStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case bufferedStream <- v:
			}
		}
	}()
	return bufferedStream
}
