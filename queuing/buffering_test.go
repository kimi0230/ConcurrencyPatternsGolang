package queuing

import (
	common "ConcurrencyPatternsGolang/utility/common"
	"bufio"
	"io"
	"testing"
)

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

/*
bufio 建立緩衝寫入
*/
func BenchmarkBufferedWrite(b *testing.B) {
	bufferredFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferredFile))
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range common.Take(done, common.Repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}
}

/*
go test -benchmem -run=none ConcurrencyPatternsGolang/queuing/ -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/queuing
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkUnbufferedWrite-8        230109              7182 ns/op               1 B/op          1 allocs/op
BenchmarkBufferedWrite-8         1000000              1652 ns/op               1 B/op          1 allocs/op
PASS
ok      ConcurrencyPatternsGolang/queuing       4.396s
*/
