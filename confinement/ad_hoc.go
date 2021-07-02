package main

import (
	"fmt"
)

/*
特定約束
可看到 loopData 和 handleData channel 上的循環都可以使用 int slice 的 data
*/
func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
