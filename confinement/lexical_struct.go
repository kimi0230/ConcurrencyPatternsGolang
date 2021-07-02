package main

import (
	"bytes"
	"fmt"
	"sync"
)

/*
可以不需要通過通訊完成內存訪問同步或共享數據

Result:
ang
gol
*/
func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3]) // <1> 前三個 elemet
	go printData(&wg, data[3:]) // <2> 第三個 elemet 之後

	wg.Wait()
}
