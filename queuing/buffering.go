package queuing

import (
	"io/ioutil"
	"log"
	"os"
)

/*
TempFile在目錄目錄中創建一個新的臨時文件，打開該文件進行讀取和寫入，並返回生成的*os.File
結束後會自動刪除它。
*/
func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return file
}
