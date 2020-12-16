package main

import "fmt"

func main() {
	/*
		recover 只有在发生 panic 之后调用才会生效。然而在上面的控制流中，recover 是在 panic 之前调用的，
		并不满足生效的条件，所以我们需要在 defer 中使用 recover 关键字。
	*/
	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")
}
