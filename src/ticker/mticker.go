package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * time.Duration(8))
	go func() {
		defer fmt.Println("defer")
		for {
			<-ticker.C
			fmt.Println(time.Now())
		}
	}()
	select {}
}
