# fan-in fan-out
fan out 啟動多個 goroutine 以處理來自 pipeline 的輸入,
fan in 將多個結果組合到一個 channel
例子是亂數產生十個質數, 並比較 pipeline 與 fan-out fan-in

## fan-in
多對一, 類似 leaky bucket. 速度取決於消費者(擋 request)

## fan-out
一對多, 類似 token bucket. 速度取決於生產者(給 redis/db 等後端資源)

## 比較結果
```shell
go test -benchmem -run=none ConcurrencyPatternsGolang/fan-out-fan-in -bench=.
```

```
goos: darwin
goarch: amd64
pkg: ConcurrencyPatternsGolang/fan-out-fan-in
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkNaivePrimeFinder-4                         1167           1057643 ns/op            2196 B/op        189 allocs/op
BenchmarkFanOutFanInNaivePrimeFinder-4              2738            398425 ns/op            2808 B/op        218 allocs/op
PASS
ok      ConcurrencyPatternsGolang/fan-out-fan-in        3.409s
```