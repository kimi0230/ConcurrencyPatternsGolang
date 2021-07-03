package fanoutfanin

import (
	"fmt"
	"testing"
)

// var tests = []struct{}{}

func TestNaivePrimeFinder(t *testing.T) {
	a := naivePrimeFinder()
	fmt.Println(a)
}

func TestFanOutFanInNaivePrimeFinder(t *testing.T) {
	a := fanOutFanInNaivePrimeFinder()
	fmt.Println(a)
}

func BenchmarkNaivePrimeFinder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		naivePrimeFinder()
	}
}

func BenchmarkFanOutFanInNaivePrimeFinder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fanOutFanInNaivePrimeFinder()
	}
}

/*
go test -benchmem -run=none ConcurrencyPatternsGolang/fan-out-fan-in -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/fan-out-fan-in
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkNaivePrimeFinder-4                         1167           1057643 ns/op            2196 B/op        189 allocs/op
BenchmarkFanOutFanInNaivePrimeFinder-4              2738            398425 ns/op            2808 B/op        218 allocs/op
PASS
ok      ConcurrencyPatternsGolang/fan-out-fan-in        3.409s
*/
