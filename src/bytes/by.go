package main

import (
	"bytes"
	"fmt"
)

func main() {

	//很多操作和strings很像，不写了
	b := []byte("learnning go编程")
	fmt.Println(string(bytes.ToUpper(b)))
	fmt.Println(b)

	//buffer 底层就是slice，当我们操作Buffer时，除了初始化和扩容时会重新申请底层内存块，其他时候只是对切片重新切片，也即只是改变了切片的len属性，以及p的指向，底层被指向的那整块内存块并不会发生改变。
	rd := bytes.NewBufferString("Hello World!")
	buf := make([]byte, 6)
	// 获取数据切片
	bb := rd.Bytes()
	// 读出一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("%s\n", rd.String()) // World!
	fmt.Printf("%s\n\n", bb)        // Hello World!

	rd.Write([]byte("abcdefg"))
	fmt.Printf("%s\n", rd.String())
	fmt.Printf("%s\n\n", b)
}
