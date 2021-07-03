package fanoutfanin

import (
	"math/rand"
)

// 重複調用function的生成器
func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
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

// take: 將傳入的 channel 的值寫入 takeStream, 取出前 num 個項目, 並回傳
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
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

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

// 回傳質數
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			prime := true
			// O(N)
			// integer -= 1
			// for divisor := integer - 1; divisor > 1; divisor-- {
			// 	if integer%divisor == 0 {
			// 		prime = false
			// 		break
			// 	}
			// }

			// 優化 O(sqrtN)
			for factor := 2; factor*factor <= integer; factor++ {
				if integer%factor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}

func naivePrimeFinder() []int {
	rand := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	// 一個 goroutine 產生亂數int
	randIntStream := toInt(done, repeatFn(done, rand))
	result := []int{}

	// 將傳入的 channel 的值寫入 takeStream, 取出前 num 個項目, 並回傳
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		result = append(result, prime.(int))
	}
	return result
}
