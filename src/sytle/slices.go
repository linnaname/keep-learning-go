package main

import "fmt"

func main() {
	var t []string
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(t == nil)

	b := []string{}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(b == nil)

	var ct map[string]string
	ct["id"] = "id"

}
