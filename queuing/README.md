# Queuing

Queue 可將stage 分離, 以便一個stage的運行時間不會影響另一個stage運行時間

Sometimes it’s useful to begin accepting work for your pipeline even though the pipeline is not yet ready for more. This process is called queuing. 
Adding queuing prematurely can hide synchronization issues such as deadlocks and livelocks, and further, as your program converges towards correctness, you may find that you need more or less queuing.
Queuing will almost never speed up the total runtime of your program; it will only allow the program to behave differently.


寫入緩衝比未寫入還快, 因為 bufio.Writer 在 writes 會 queue至內部的buffer直到chunk塞滿, 然後寫出, 稱之為 chunking

```shell
go test -benchmem -run=none ConcurrencyPatternsGolang/queuing/ -bench=.
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/queuing
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkUnbufferedWrite-8        230109              7182 ns/op               1 B/op          1 allocs/op
BenchmarkBufferedWrite-8         1000000              1652 ns/op               1 B/op          1 allocs/op
PASS
ok      ConcurrencyPatternsGolang/queuing       4.396s
```