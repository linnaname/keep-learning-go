package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := "foo   bar      baz"
	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
