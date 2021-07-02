# Preventing Goroutine Leaks

## [ main goroutine done ](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/main_goroutine_done.go)
main 的 goroutine 先結束, 導致子 goroutine 無法被 GC
解法: 使用 context, 通知結束
[solution](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/main_goroutine_done_solution.go)

## [ nil channel ](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/nil_channel.go)
收到 nil 的 channel 阻塞
解法: 將父子 goroutine之間建立一個`信號通道`, 讓父 goroutine可以向子發出取消訊號
[solution](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/nil_channel_solution.go)

## [ blocked channel write ](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/blocked_channel_write.go)
一個 goroutine 阻塞了向 channel 進行寫入的請求. 與第一個範例相似
解法: 為生產者 goroutine 提供一個通知他退出的 channel

[solution](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks/blocked_channel_write_solution.go)