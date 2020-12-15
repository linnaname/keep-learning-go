package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./linnana.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := io.WriteString(file, "Go is fun!"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
