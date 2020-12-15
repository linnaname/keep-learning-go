package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	sl := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	var buf bytes.Buffer
	for _, str := range sl {
		n, err := buf.Write([]byte(str))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if n != len(str) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println(buf.String())
}
