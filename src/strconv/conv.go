package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "999"
	num, _ := strconv.Atoi(numStr)
	fmt.Println(num)
	fmt.Println(strconv.Itoa(num))

	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("f"))

	strF := "250.56"
	str, err := strconv.ParseFloat(strF, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type:%T value:%#v\n", str, str)

	numF := 250.56
	strFF := strconv.FormatFloat(numF, 'f', 4, 64)
	fmt.Printf("type:%T value:%#v\n", strFF, strFF)
	fmt.Println("This is", strconv.Quote("a"))

}
