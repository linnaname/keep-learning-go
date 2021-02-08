package main

import (
	"fmt"
	"time"
)

func main() {

	select {
	case <-time.After(time.Minute * time.Duration(1)):
		fmt.Println("time out")
		return
	}

	fmt.Println("return finally")

}
