package main

import (
	"fmt"
	"math/rand"
)

/*
Result:
3 random ints:
1: 5577006791947779410
2: 8674665223082153551
3: 6129484611666145821
*/
func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // <1> 沒執行到這裡
			defer close(randStream)
			for {
				randStream <- rand.Int() // 阻塞, 寫不進去. 因為main routine只讀了3次就離開了
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
