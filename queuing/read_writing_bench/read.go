package readwritingbench

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 讀文件方式一：利用ioutil.ReadFile直接從文件讀取到[]byte中
func readReadFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

// 讀文件方式二：先從文件讀取到file中，在從file讀取到buf, buf在追加到最終的[]byte
func readByteBuf(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

// 讀文件方式三：先從文件讀取到file, 在從file讀取到Reader中，從Reader讀取到buf, buf最終追加到[]byte
func readBufioNewReader(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

// 讀文件方式四：讀取到file中，再利用ioutil將file直接讀取到[]byte中
// 最佳
func readReadAll(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)
	return string(fd)
}
