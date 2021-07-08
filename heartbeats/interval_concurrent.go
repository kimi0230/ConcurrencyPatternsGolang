package main

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
)

func DoWork(done <-chan interface{}, pulseInterval time.Duration, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(pulseInterval)
	numLoop: // <2>
		for _, n := range nums {
			for { // <1>
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					time.Sleep(500 * time.Millisecond)
					continue numLoop // <3>
				}
			}
		}
	}()

	return heartbeat, intStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4}
	const timeout = 2 * time.Second
	heartbeat, results := DoWork(done, timeout/4, intSlice...)

	<-heartbeat // <4>

	for {
		select {
		case r, ok := <-results:
			if ok == false {
				fmt.Println("Done")
				return
			} else {
				fmt.Println(r)
			}
		case <-heartbeat: // <5>
			fmt.Println("pulse")
		case <-time.After(timeout):
			log.Fatal("test timed out")
		}
	}
}
