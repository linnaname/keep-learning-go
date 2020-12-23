package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := command("linnaname")
	if err != nil {
		fmt.Println("command error: ", err)
		return
	}
	fmt.Println(output)
}

func command(name string) (string, error) {
	out, err := exec.Command("/usr/bin/java", "HelloWorld", name).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
