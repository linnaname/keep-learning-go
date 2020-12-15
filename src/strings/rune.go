package main

import "fmt"

func main() {
	/**
	Rune 是int32 的别名。用UTF-8 进行编码。这个类型在什么时候使用呢？例如需要遍历字符串中的字符。可以循环每个字节（仅在使用US ASCII 编码字符串时与字符等价，而它们在Go中不存在！）。因此为了获得实际的字符，需要使用rune类型。在UTF-8 世界的字符有时被称作runes。通常，当人们讨论字符时，多数是指8 位字符。UTF-8 字符可能会有32 位，称作rune。
	*/
	s := "Go编程"

	fmt.Println(len(s))
	fmt.Println(len(string(rune('编'))))
	fmt.Println(len([]rune(s)))
}
