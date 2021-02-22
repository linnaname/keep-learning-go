package main

import "fmt"

func main() {
	var pi *int = nil
	fmt.Println(pi == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = pi
	fmt.Println(i == nil)
}
