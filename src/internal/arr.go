package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}

	fmt.Printf("%T", arr1)
	fmt.Printf("%T", arr2)

	fmt.Println(reflect.DeepEqual(arr1, arr2))
}
