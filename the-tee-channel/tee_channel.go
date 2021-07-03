package theteechannel

import (
	theordonechannel "ConcurrencyPatternsGolang/the-or-done-channel"
)

func Tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)

		for val := range theordonechannel.OrDone(done, in) {
			var out1, out2 = out1, out2 // 將要使用的 out1 out2 變成private 變數
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil // 設成 nil 來阻塞, 而另一個 channel可繼續
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
