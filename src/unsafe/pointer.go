package main

import (
	"fmt"
	"unsafe"
)

/**
指针例子
*/

func main() {
	a := int(100)
	b := int(100)
	fmt.Println(a == b)

	//这是不允许的
	//c := int(100)
	//var d *float64
	//d = &c

	/**
		slice的结构
		type slice struct {
	    array unsafe.Pointer // 元素指针
	    len   int // 长度
	    cap   int // 容量
		}
	*/

	//指针运算
	/**
	Len，cap 的转换流程如下：
	Len: &s => pointer => uintptr => pointer => *int => int
	Cap: &s => pointer => uintptr => pointer => *int => int
	*/
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20

}
