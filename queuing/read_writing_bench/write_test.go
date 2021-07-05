package readwritingbench

import "testing"

var strTest string = "測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試__"

// 方式一：使用 io.WriteString
func BenchmarkWriteWriteString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WriteWriteString("file/test1", strTest)
	}
}

// 方式二：使用 ioutil.WriteFile
func BenchmarkWriteWriteFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WriteWriteFile("file/test2", strTest)
	}
}

// 方式三：使用 File(Write, byte)
func BenchmarkWriteByte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WriteByte("file/test3", strTest)
	}
}

// 方式四：使用 bufio.NewWriter
func BenchmarkWriteNewWriter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WriteNewWriter("file/test4", strTest)
	}
}

/*
go test -benchmem -run=none ConcurrencyPatternsGolang/queuing/read_writing_bench -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/queuing/read_writing_bench
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkWriteWriteString-8                 7966            136062 ns/op             824 B/op          4 allocs/op
BenchmarkWriteWriteFile-8                   8446            135213 ns/op             824 B/op          4 allocs/op
BenchmarkWriteByte-8                        8271            138011 ns/op             824 B/op          4 allocs/op
BenchmarkWriteNewWriter-8                   7939            137542 ns/op            4216 B/op          4 allocs/op
*/
