package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
		}
		fmt.Println(n, string(buf[:n]))
	}
}
