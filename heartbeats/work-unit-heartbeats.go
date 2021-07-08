package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DoWork(done <-chan interface{}) (<-chan interface{}, <-chan int) {
	// This ensures that there’s always at least one pulse sent out even if no one is listening in time for the send to occur.
	heartbeatStream := make(chan interface{}, 1)
	workStream := make(chan int)
	go func() {
		defer close(heartbeatStream)
		defer close(workStream)

		for i := 0; i < 10; i++ {
			// result 和 heartbeats 切開
			// 如果接收者沒有準備好接收結果, 作為替代將會收到一個heartbeat
			select {
			case heartbeatStream <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case workStream <- rand.Intn(10):
			}
		}
	}()

	return heartbeatStream, workStream
}

func main() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) })

	heartbeat, results := DoWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}

/*
we receive one pulse for every result, as intended.

pulse
results 1
pulse
results 7
pulse
results 7
pulse
results 9
pulse
results 1
pulse
results 8
pulse
results 5
pulse
results 0
pulse
results 6
pulse
results 0

*/
