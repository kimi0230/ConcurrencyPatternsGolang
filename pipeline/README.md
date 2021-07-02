# pipeline
pipeline 當你的程序需要流動式或批次處理時, 可以相互獨立的修改每個 stage. 
混合搭配 stage 組出一個新的數據流, 而無需修改 stage

## [additional stage to pipeline](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/pipeline/additional-stage-to-pipeline.go)
分離每個 stage, 一步一步將每個 function 回傳值帶入下一個 function.

## [func stream processing](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/pipeline/func-stream-processing.go)
修改上一個範例, 避免每stage 都創建一個新的slice. 但是不得不將pipeline 寫入到 for裏面,
不僅限制重複利用pipeline, 也限制了擴充性

## 最佳實現方式: [chan stream processing](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/pipeline/chan-stream-processing.go)
channel 非常適合在 go 中建立 pipeline

```shell
go test -benchmem -run=none ConcurrencyPatternsGolang/pipeline/some-handy-generators/pipelines_test/ -bench=.
```
綁定類型的 stage 速度是 interface 類型的兩倍.
如果一個 state 在計算上花費很大, 可以藉由 [fan-out-fan-in](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/fan-out-fan-in) 來緩解這問題.