package main

import (
	"fmt"
	"unsafe"
)

type Cat struct {
	color string
	age   int32
	name  string
}

func main() {
	cat := Cat{
		color: "red",
		age:   1,
		name:  "lin",
	}
	fmt.Println(unsafe.Sizeof(cat))
	fmt.Println(unsafe.Offsetof(cat.name))
	//Alignof 返回 m，m 是指当类型进行内存对齐时，它分配到的内存地址能整除 m
	fmt.Println(unsafe.Alignof(cat.age))
}
