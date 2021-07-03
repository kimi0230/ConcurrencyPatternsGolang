# the or done channel
有時候你需要處理來自系統各個分散部分的channel. 
與操作pipeline所不同的是你不能通過 done channel 進行取消, 
你不知道你的goroutine 是否被取消
所以需要用 select 將 read 包裝至裏面

we need to wrap our read from the channel with a select statement that also selects from a done channel. This is perfectly fine, but doing so takes code that’s easily read like this:
