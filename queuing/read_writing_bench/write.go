package readwritingbench

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 方式一：使用 io.WriteString
func WriteWriteString(path, str string) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		fmt.Println("create file fail")
		return
	}

	//將文件寫入
	_, writeErr := io.WriteString(f, str)
	if writeErr != nil {
		fmt.Println("write fail")
		return
	}
}

// 方式二：使用 ioutil.WriteFile
func WriteWriteFile(path, str string) {
	var d = []byte(str)
	err := ioutil.WriteFile(path, d, 0666)
	if err != nil {
		fmt.Println("write fail")
		return
	}
}

// 方式三：使用 File(Write, byte)
func WriteByte(path, str string) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		fmt.Println("create file fail")
		return
	}

	var d1 = []byte(str)
	_, writeErr := f.Write(d1)
	if writeErr != nil {
		fmt.Println("write fail")
		return
	}
	// f.Sync()
}

// 方式四：使用 bufio.NewWriter
func WriteNewWriter(path, str string) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		fmt.Println("create file fail")
	}

	w := bufio.NewWriter(f) //創建新的 Writer 對象
	_, writeErr := w.WriteString(str)
	if writeErr != nil {
		fmt.Println("write fail")
		return
	}
	w.Flush()
}
