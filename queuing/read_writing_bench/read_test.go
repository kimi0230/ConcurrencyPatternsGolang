package readwritingbench

import "testing"

const path = "test.log"

func BenchmarkReadReadFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readReadFile(path)
	}
}

func BenchmarkReadByteBuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readByteBuf(path)
	}
}

func BenchmarkReadBufioNewReader(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readBufioNewReader(path)
	}
}

func BenchmarkReadReadAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readReadAll(path)
	}
}

/*
go test -benchmem -run=none ConcurrencyPatternsGolang/queuing/read_writing_bench -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/queuing/read_writing_bench
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkReadReadFile-8                    47835             29645 ns/op           27464 B/op          6 allocs/op
BenchmarkReadByteBuf-8                     19966             54775 ns/op           82040 B/op         12 allocs/op
BenchmarkReadBufioNewReader-8              21823             56503 ns/op           87160 B/op         14 allocs/op
BenchmarkReadReadAll-8                     24049             57199 ns/op           78072 B/op         16 allocs/op
PASS
ok      ConcurrencyPatternsGolang/queuing/read_writing_bench    7.349s
*/
