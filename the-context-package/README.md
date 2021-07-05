# The Context Package
使用 context 在不同 Goroutine 之間同步請求特定數據, 取消訊號, deadline

in concurrent programs it’s often necessary to preempt operations because of timeouts, cancellation, or failure of another portion of the system. We’ve looked at the idiom of creating a [done](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/the-or-done-channel) channel which flows through your program and cancels all blocking concurrent operations.

[Preventing goroutine leaks](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/preventing-goroutine-leaks), cancellation in a function has three aspects:

* A goroutine’s parent may want to cancel it
* A goroutine may want to cancel its children
* Any blocking operations within a goroutine need to be preemptable so that it may be canceled.

The context package helps manage all three of these.

context 的 key 和 val 都是被定義為 interface{}
Use context values only for request-scoped data that transits processess and API boundaries, not for passing optional paramenters.
