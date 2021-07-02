package main

import (
	"fmt"
)

/*
只將 channel 的讀或寫暴露給需要的 goroutine

Result:
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Done receiving!
*/
func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // <1> 防止其他人寫入
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // <3> 只讀
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // <2>
	consumer(results)
}
