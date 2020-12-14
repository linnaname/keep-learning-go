package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	handleFunc := func() {
		fmt.Println("Hello Once")
	}

	for i := 0; i < 20; i++ {
		go func() {
			once.Do(handleFunc)
		}()
	}

	time.Sleep(10 * time.Second)
}
