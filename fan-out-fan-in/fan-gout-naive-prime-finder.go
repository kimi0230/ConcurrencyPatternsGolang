package fanoutfanin

import (
	"math/rand"
	"runtime"
	"sync"
)

// 將多個 channel 合併成 一個channel 並回傳
func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} { // <1>
	var wg sync.WaitGroup // <2>
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) { // <3>
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// Select from all the channels
	wg.Add(len(channels)) // <4>
	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete
	go func() { // <5>
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func fanOutFanInNaivePrimeFinder() []int {

	done := make(chan interface{})
	defer close(done)

	// start := time.Now()

	rand := func() interface{} { return rand.Intn(50000000) }

	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)

	// fan out 產生多個 primeFinder的goroutine
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	result := []int{}
	for prime := range take(done, fanIn(done, finders...), 10) {
		result = append(result, prime.(int))
	}
	return result
}
