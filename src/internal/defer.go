package main

import (
	"fmt"
	"time"
)

func main() {
	startedAt := time.Now()
	//当这个defer调用时，startedAt就已经被拷贝了，应该通过defer func实现
	defer fmt.Println(time.Since(startedAt))

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")

	time.Sleep(time.Second)
}
