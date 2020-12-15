package main

import (
	"errors"
	"fmt"
	"io"
)

type MyReader struct {
	src string
	cur int
}

func New(src string) *MyReader {
	return &MyReader{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (mr *MyReader) Read(p []byte) (n int, err error) {
	if mr.cur >= len(mr.src) {
		return 0, io.EOF
	}

	// x 是剩余未读取的长度
	x := len(mr.src) - mr.cur
	n, bound := 0, 0
	if x >= len(p) {
		// 剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		// 每次读取一个字节，执行过滤函数
		if char := alpha(mr.src[mr.cur]); char != 0 {
			buf[n] = char
		}
		n++
		mr.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil
}

func main() {
	reader := New("Hello! It's 9am, where is the sun?")
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
