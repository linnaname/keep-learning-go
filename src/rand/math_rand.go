package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(newId())
}

func newId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
