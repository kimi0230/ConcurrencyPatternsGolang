package workerqueue

import (
	"fmt"
	"sync"
	"time"
)

var concurrencyProcesses = 5

func WorkQueue(jobs []string, limit int) {
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	queue := make(chan string)

	// 第一步驟建立 queue 通道，並將所有的內容都丟進 queue 內
	go func(queue chan<- string) {
		for i := 0; i < len(jobs); i++ {
			queue <- jobs[i]
		}
		close(queue)
	}(queue)

	if limit > 0 {
		concurrencyProcesses = limit
	}

	// 第二步 建立 concurrencyProcesses 個 goroutine
	for i := 0; i < concurrencyProcesses; i++ {
		go func(queue <-chan string) {
			for job := range queue {
				func() {
					defer wg.Done()
					fmt.Println("Working :", job)
					time.Sleep(1 * time.Millisecond)
				}()
			}
		}(queue)
	}
	wg.Wait()
}
