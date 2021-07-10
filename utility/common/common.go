package common

// Repeat: 將你傳入的值傳入channel, 直到你告訴他停止
func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

// RepeatFn : 重複調用function的生成器
func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

// Take: 將傳入的 channel 的值寫入 takeStream, 取出前 num 個項目, 並回傳
func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func Or(channels ...<-chan interface{}) <-chan interface{} {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	switch len(channels) {
	case 0: // <2>
		return nil
	case 1: // <3>
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() { // <4>
		defer close(orDone)

		switch len(channels) {
		case 2: // <5>
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // <6>
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...): // <6>
			}
		}
	}()
	return orDone
}
