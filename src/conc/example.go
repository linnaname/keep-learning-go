package main

import (
	"fmt"
	"time"
)

func main() {
	var running int32 = 1

	go func() {
		fmt.Println("start routine1")
		for running == 1 {
			fmt.Printf("routine1 runnning %d\n", running)
		}
		fmt.Printf("routine1 end %d\n", running)
	}()

	go func() {
		fmt.Println("start routine2")
		for {
			running = 0
		}
	}()
	time.Sleep(time.Minute * 2)
}
