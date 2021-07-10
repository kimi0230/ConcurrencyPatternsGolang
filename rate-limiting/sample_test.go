package rateLimiter

import (
	"testing"
)

func TestNoRateLimit(t *testing.T) {
	api := OpenNoRateLimit()
	api.conn = api
	api.DemoFunc()
}

func TestRateLimit(t *testing.T) {
	api := OpenRateLimit()
	api.conn = api
	api.DemoFunc()
}

func TestMultiRateLimit(t *testing.T) {
	api := OpenMultiRateLimit()
	api.conn = api
	api.DemoFunc()
}

func BenchmarkNoRateLimit(b *testing.B) {
	b.ResetTimer()
	api := OpenNoRateLimit()
	api.conn = api
	for i := 0; i < b.N; i++ {
		api.DemoFunc()
	}
}

func BenchmarkRateLimit(b *testing.B) {
	b.ResetTimer()
	api := OpenRateLimit()
	api.conn = api
	for i := 0; i < b.N; i++ {
		api.DemoFunc()
	}
}

/*
go test -benchmem -run=none ConcurrencyPatternsGolang/rate-limiting/ -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/rate-limiting
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkNoRateLimit-4              2328            525253 ns/op            3711 B/op         82 allocs/op
BenchmarkRateLimit-4                2474            486217 ns/op           10790 B/op        188 allocs/op
PASS
ok      ConcurrencyPatternsGolang/rate-limiting 2.538s
*/
