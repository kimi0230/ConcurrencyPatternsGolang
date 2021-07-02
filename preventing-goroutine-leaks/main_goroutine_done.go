package main

import "fmt"

func main() {
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break // 導致goroutine處於無法被回收的狀態, 可以用context來避免這問題
		}
	}
}
