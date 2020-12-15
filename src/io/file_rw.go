package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	sl := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	file, err := os.Create("./linnana.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, str := range sl {
		n, err := file.Write([]byte(str))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(str) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")

	fileR, err := os.Open("./linnana.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileR.Close()

	by := make([]byte, 4)
	for {
		n, err := fileR.Read(by)
		if err == io.EOF {
			break
		}
		fmt.Print(string(by[:n]))
	}
	fmt.Println("file read done")

	bytes, err := ioutil.ReadFile("./linnana.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)

}
