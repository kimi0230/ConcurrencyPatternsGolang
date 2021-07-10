package main

import (
	"ConcurrencyPatternsGolang/utility/common"
	"log"
	"os"
	"time"
)

type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration,
) (heartbeat <-chan interface{}) // <1> 定義監控和重啟的 goroutine 訊號

// 一個管理員監控 goroutine , timeout 變數, 一個 startGoroutine 來請動他的監控的 goroutine
// 也返回一個 startGoroutineFn, 表示管理員本身也是可以監控的
func newSteward(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn { // <2>
	return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() { // <3> 啟動監控
				wardDone = make(chan interface{})                                    // <4>
				wardHeartbeat = startGoroutine(common.Or(wardDone, done), timeout/2) // <5>
			}
			startWard()
			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)

				for { // <6> 確保可以收到心跳
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat: // <7> 確認管理員可以發出自己的心跳
						continue monitorLoop
					case <-timeoutSignal: // <8> 沒收到管理員心跳 -> timeout
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard() // 從新啟動一個goroutine 繼續監控
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()

		return heartbeat
	}
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I'm irresponsible!")
		go func() {
			<-done // <1> 只是等待被取消, 沒發出任何心跳
			log.Println("ward: I am halting.")
		}()
		return nil
	}
	doWorkWithSteward := newSteward(4*time.Second, doWork) // <2> 設置doWork的超時時間  4s

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkWithSteward(done, 4*time.Second) { // <4> 啟動管理員
	}
	log.Println("Done")
}
