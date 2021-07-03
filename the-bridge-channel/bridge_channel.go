package thebridgechannel

import theordonechannel "ConcurrencyPatternsGolang/the-or-done-channel"

func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{}) // <1>
	go func() {
		defer close(valStream)
		for { // <2>
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range theordonechannel.OrDone(done, stream) { // <3>
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
