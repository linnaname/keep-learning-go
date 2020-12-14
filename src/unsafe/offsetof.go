package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func main() {
	p := Programmer{"test", "go"}

	fmt.Println(p)

	/*
		Offsetof 获取成员偏移量,做地址运算后可以直接改变对应的值
	*/
	//结构体会被分配一块连续的内存，结构体的地址也代表了第一个成员的地址。
	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)

	//unsafe.Sizeof返回的是类型数据结构的大小而不是其指向内容的数据大小
	newLang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(string(""))))
	*newLang = "Java"

	fmt.Println(p)

	fmt.Println(unsafe.Sizeof(string("中文")))
	fmt.Println(unsafe.Sizeof(string("CN")))
	fmt.Println(unsafe.Sizeof(string("qcraoaaaaa")))
	fmt.Println(unsafe.Sizeof(string("")))
	fmt.Println(unsafe.Sizeof(int(0)))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(unsafe.Sizeof(slice)) //24

}
