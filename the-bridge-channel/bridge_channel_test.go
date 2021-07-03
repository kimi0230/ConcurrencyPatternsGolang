package thebridgechannel

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	// 建立 10 個 channel 每個 channel都寫入一個元素, 並將這些 channel 傳入橋接函數
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range Bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
