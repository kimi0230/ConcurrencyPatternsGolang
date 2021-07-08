# Heartbeats

Because of the heartbeat, we can safely write our test without timeouts. The only risk we run is of one of our iterations taking an inordinate amount of time. If that’s important to us, we can utilize the safer interval-based heartbeats and achieve perfect safety.

* [一段時間間隔發出心跳](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/heartbeats/interval-heartbeats.go)
* [一個工作單位發出心跳](https://github.com/kimi0230/ConcurrencyPatternsGolang/tree/master/heartbeats/work-unit-heartbeats.go)