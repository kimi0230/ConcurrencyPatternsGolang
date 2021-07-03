# The bridge channel
在某些情況, 你可能會想要從一連串的 channels 中取得值
In some circumstances, you may find yourself wanting to consume values from a sequence of channels:
<-chan <-chan interface{}

通過橋接, 我們可以用一個 for loop 來處理 channel的channel, 並專注於我們的 loop 邏輯. 將傳遞的channel拆解成單一個channel來傳遞.來方便處理邏輯

Thanks to bridge, we can use the channel of channels from within a single range statement and focus on our loop’s logic. Destructuring the channel of channels is left to code that is specific to this concern.