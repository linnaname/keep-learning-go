package main

import (
	"fmt"
	"runtime"
)

func main() {
	//2,3代表底层的调用位置
	for i := 0; i < 4; i++ {
		//1代表这个位置
		call(i)
	}

	//函数把当前go程调用栈上的调用栈标识符填入切片pc中，返回写入到pc中的项数
	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs)
	fmt.Println(pcs[:i])

	for _, pc := range pcs[:i] {
		funcPC := runtime.FuncForPC(pc)
		file, line := funcPC.FileLine(pc)
		println(funcPC.Name(), file, line)
	}
}

func call(skip int) {
	//0代表发起Caller的位置
	pc, file, line, ok := runtime.Caller(skip)
	//获取函数名
	pcName := runtime.FuncForPC(pc).Name()
	fmt.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
}
