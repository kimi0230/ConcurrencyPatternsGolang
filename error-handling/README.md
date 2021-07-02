# error handling

## [imporoper err handling](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/error-handling/imporoper-err-handling.go)
不能簡單的把錯誤回傳
解決: 新增了一個 struct, 當作回傳的 channel, 裡面記錄 error
```go
type Result struct { 
    Error    error
    Response *http.Response
}
```
[solution](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/proper-err-handling.go)