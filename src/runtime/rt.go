package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	//注意这个会STW，所以设置应该尽早进行
	runtime.GOMAXPROCS(1)
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	fmt.Println("version:", runtime.Version())

	go say("world")
	say("hello")
	//从注释来看这个GC是立即执行的，和Java的并不相同
	runtime.GC()

	debug.SetGCPercent(100)
	debug.SetMaxStack(1000)
	debug.SetMaxThreads(100)
}

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
