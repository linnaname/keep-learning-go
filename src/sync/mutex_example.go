package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	//在没有加锁前可以使用race工具检测
	var locker sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				locker.Lock()
				count++
				locker.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)

}
