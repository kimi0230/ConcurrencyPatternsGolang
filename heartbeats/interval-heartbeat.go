package main

import (
	"fmt"
	"time"
)

func DoWork(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)
	go func() {
		defer close(heartbeat)
		defer close(results)
		pulse := time.Tick(pulseInterval)       // 心跳間隔
		workGen := time.Tick(2 * pulseInterval) // 模擬心跳行為的channel, 這樣能看到從goroutine中發出的一些心跳

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default: // 可能沒有人接收得到heartbeat
			}
		}
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse: // 就像 done channel 一樣, 當你執行或發送時, 需要一個處理併發送心跳的case
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen: // 間隔兩秒發送一次
				sendResult(r)
			}
		}
	}()
	return heartbeat, results
}

func main() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) })

	const timeout = 2 * time.Second
	heartbeat, results := DoWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results: // <5>
			if !ok {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout): // 超時
			return
		}
	}
}
