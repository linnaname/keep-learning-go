package main

import (
	"fmt"
	"sync"
)

func main() {
	var locker sync.RWMutex
	locker.RLock()
	fmt.Println("read lock 1 success")
	locker.RLock()
	fmt.Println("read lock 2 sucess")
	locker.RUnlock()
	locker.RUnlock()
}
