package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

const (
	capacity    = 10
	consumerNum = 3
	producerNum = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	product := make(chan int, capacity)
	producer(product)
	consumer(product)
	<-quit
}

func producer(out chan<- int) {
	for i := 0; i < producerNum; i++ {
		go func(nu int) {
			for {
				cond.L.Lock()
				for len(out) == capacity {
					fmt.Println("Capacity Full, stop Produce")
					cond.Wait()
				}
				num := rand.Intn(100)
				out <- num
				fmt.Printf("Produce %d produce: num %d\n", nu, num)
				cond.L.Unlock()
				cond.Signal()
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func consumer(in <-chan int) {
	for i := 0; i < consumerNum; i++ {
		go func(nu int) {

			for {
				cond.L.Lock()
				for len(in) == 0 {
					fmt.Println("Capacity Empty, stop Consume")
					cond.Wait()
				}
				num := <-in
				fmt.Printf("Goroutine %d: consume num %d\n", nu, num)
				cond.L.Unlock()
				time.Sleep(time.Millisecond * 500)
				cond.Signal()
			}
		}(i)
	}
}
