package sample

import (
	"testing"
)

func BenchmarkNoRateLimit(b *testing.B) {
	b.ResetTimer()
	api := OpenNoRateLimit()
	for i := 0; i < b.N; i++ {
		api.DemoFunc()
	}
}

func BenchmarkRateLimit(b *testing.B) {
	b.ResetTimer()
	api := OpenRateLimit()
	for i := 0; i < b.N; i++ {
		api.DemoFunc()
	}
}

/*
go test -benchmem -run=none ConcurrencyPatternsGolang/rate-limiting/sample/ -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/rate-limiting/sample
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkNoRateLimit-4             63008             21317 ns/op             264 B/op         22 allocs/op
BenchmarkRateLimit-4               64429             19500 ns/op             264 B/op         22 allocs/op
PASS
ok      ConcurrencyPatternsGolang/rate-limiting/sample  2.993s
*/
