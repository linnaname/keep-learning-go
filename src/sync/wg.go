package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("1")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("2")
	}()
	wg.Wait()
	fmt.Println("finished")
}
