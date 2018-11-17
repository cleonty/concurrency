package main

func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for num := range intStream {
			if isPrime(num) {
				select {
				case <-done:
					return
				case primeStream <- num:
				}
			}
		}
	}()
	return primeStream
}

func isPrime(num int) bool {
	for integer := num - 1; integer > 1; integer-- {
		if num%integer == 0 {
			return false
		}
	}
	return true
}
