package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	locker := &sync.Mutex{}
	cond := sync.NewCond(locker)
	for i := 0; i < 20; i++ {
		handle(cond, i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("===================")
	cond.Signal()
	time.Sleep(2 * time.Second)
	fmt.Println("===================")
	cond.Broadcast()
	time.Sleep(10 * time.Second)
}

func handle(cond *sync.Cond, i int) {
	go func(cond *sync.Cond, i int) {
		cond.L.Lock()
		for condition() {
			fmt.Println("-goroutine-" + strconv.Itoa(i) + " 命中wait")
			cond.Wait()
		}
		fmt.Println("-goroutine-" + strconv.Itoa(i) + " 命中条件")
		cond.L.Unlock()
	}(cond, i)
}

func condition() bool {
	rand.Intn(50)
	if rand.Intn(50) > 20 {
		fmt.Print(true)
		return true
	}
	fmt.Print(false)
	return false
}
