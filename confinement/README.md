# Confinement (約束)

## [ad hoc (特定約束) ](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/confinement/ad_hoc.go)
特定約束是指通過公約協議實現約束, 很難在任何規模的項目上進行協調,
除非你有工具在每次有人提交代碼時對你的代碼進行靜態分析

## [lexical (詞法約束) ](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/confinement/lexical.go)
詞法約束, 作用域僅公開於特定的 goroutine
Lexical confinement involves using lexical scope to expose only the correct data and concurrency primitives for multiple concurrent processes to use

## [lexical struct ](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/confinement/lexical_struct.go)
透過傳遞不同的 slice 的 subsets, 因此約束了 goroutine 開始的 slice.
所以不需要通過通訊來完成同步內存或共享數據

如果我們有同步功能, 為什麼要約束? 可提高性能並降低開發元的認知負擔.
利用詞法約束(lexical)的併發代碼通常比不具有詞法約束的併發代碼更好理解
