package main

import "fmt"

func main() {
	testDeadLoop()
	testPointer()
	clearSL()
}

func testPointer() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	//正确的做法应该是使用 &arr[i] 替代 &v
	for _, v := range arr {
		fmt.Println(v)
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func testDeadLoop() {
	arr := []int{1, 2, 3}
	//为什么这个不会死循环呢？哈哈
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func clearSL() {
	arr := []int{1, 2, 3}
	for i, _ := range arr {
		arr[i] = 0
	}
}
