package theordonechannel

import (
	"fmt"
	"testing"
	"time"
)

var tests = []struct {
	arg1 chan interface{}
}{
	{
		make(chan interface{}),
	},
}

func TestOrDone(t *testing.T) {
	for _, tt := range tests {
		done := make(chan interface{})
		// 模擬丟值到 channel
		go func() {
			defer close(tt.arg1)
			for i := 0; i < 10; i++ {
				select {
				case <-done:
					return
				case tt.arg1 <- i:
				}
				time.Sleep(200 * time.Millisecond)
			}
		}()
		// 模擬發出 done
		go func() {
			for {
				select {
				case <-done:
					return
				case <-time.After(time.Second * 1):
					fmt.Print("timeout")
					close(done)
				}
			}
		}()

		for val := range OrDone(done, tt.arg1) {
			fmt.Println(val)
		}
	}
}
