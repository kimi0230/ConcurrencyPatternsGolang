package main

import (
	thebridgechannel "ConcurrencyPatternsGolang/the-bridge-channel"
	"ConcurrencyPatternsGolang/utility/common"
	"fmt"
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
	doWorkFn := func(done <-chan interface{}, intList ...int) (startGoroutineFn, <-chan interface{}) { // <1>
		intChanStream := make(chan (<-chan interface{})) // <2>
		intStream := thebridgechannel.Bridge(done, intChanStream)

		doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} { // <3>
			intStream := make(chan interface{}) // <4>
			heartbeat := make(chan interface{})
			go func() {
				defer close(intStream)
				select {
				case intChanStream <- intStream: // <5> 將此 channel的值丟給 bridge
				case <-done:
					return
				}

				pulse := time.Tick(pulseInterval)

				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Printf("negative value: %v\n", intVal) // <6> 收到錯誤, 離開 goroutine
							return
						}

						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}
						}
					}
				}
			}()
			return heartbeat
		}
		return doWork, intStream
	}

	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)      // <1>
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork) // <2>
	doWorkWithSteward(done, 1*time.Hour)                        // <3>

	for intVal := range common.Take(done, intStream, 6) { // <4>
		fmt.Printf("Received: %v\n", intVal)
	}
}
