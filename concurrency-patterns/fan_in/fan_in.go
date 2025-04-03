package main

import "sync"

func fanIn(channels ...chan string) chan string {
	var wg sync.WaitGroup
	multiplexedStream := make(chan string)
	multiplex := func(c <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for n := range c {
			multiplexedStream <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c, &wg)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
