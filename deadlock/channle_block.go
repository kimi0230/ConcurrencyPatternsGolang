package main

import "fmt"

func block() {
	c := make(chan int)
	c <- 118 // 會一直阻塞, 直到出現消費者. 無容量的chan 是同步的
	fmt.Println(<-c)
}

// 使用容量的channel
func solution1() {
	c := make(chan int, 1)
	c <- 118
	fmt.Println(<-c)
}

func solution2() {
	c := make(chan int)
	go func() {
		c <- 118
	}()

	fmt.Println(<-c)
}

func main() {
	// block()
	solution1()
	solution2()
}
