package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type MyReaderExt struct {
	//这样接可以复用所有实现了Reader接口的能力
	reader io.Reader
}

func NewReader(reader io.Reader) *MyReaderExt {
	return &MyReaderExt{reader: reader}
}

func filter(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (mr *MyReaderExt) Read(p []byte) (int, error) {
	n, err := mr.reader.Read(p)
	if err != nil {
		return n, err
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := filter(p[i]); char != 0 {
			buf[i] = char
		}
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil
}

func main() {
	reader := NewReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if errors.Is(err, io.EOF) {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
