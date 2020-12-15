package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	/**
	readSlice它返回的 []byte 是指向 Reader 中的 buffer，而不是 copy 一份返回，也正因为如此，通常我们会使用 ReadBytes 或 ReadString。
	很显然，ReadBytes 返回的 []byte 不会是指向 Reader 中的 buffer
	*/
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)

	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))

	r := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	l1, _ := r.ReadBytes('\n')
	fmt.Printf("the line:%s\n", l1)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n1, _ := r.ReadBytes('\n')
	fmt.Printf("the line:%s\n", l1)
	fmt.Println(string(n1))
}
