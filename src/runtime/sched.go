package main

import (
	"fmt"
	"runtime"
)

func main() {
	go print()
	runtime.Gosched()
	fmt.Println("继续执行")
}
func print() {
	fmt.Println("执行打印方法")
}
