package main

import (
	"errors"
	"fmt"
	"runtime"
)

type stack []uintptr

type mError struct {
	msg string
	*stack
}

func main() {
	nerr := errors.New("test")
	conent, err := openFile()
	fmt.Println(errors.Is(err, nerr))
	fmt.Println(errors.Unwrap(nerr))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(conent))
	}
}

func (me *mError) Error() string {
	return me.msg
}

func New(text string) error {
	return &mError{
		msg:   text,
		stack: callers(),
	}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func openFile() ([]byte, error) {
	return nil, New("my little error")
}
