package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 第一個已經lock a 然後要去lock b
	// 第二個已經lock b 然後要去lock a
	type value struct {
		mu    sync.Mutex
		value int
	}

	wg := sync.WaitGroup{}
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Println("sum = %v \n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
