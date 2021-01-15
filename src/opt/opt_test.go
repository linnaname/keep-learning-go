package main

import (
	"strings"
	"testing"
)

var gogogo = strings.Repeat("Go", 1024)

func f() {
	for range []byte(gogogo) {
	}
}

func g() {
	bs := []byte(gogogo)
	for range bs {
	}
}

func TestT1(t *testing.T) {
	t.Log(testing.AllocsPerRun(2, f)) // 0
	t.Log(testing.AllocsPerRun(2, g)) // 1
}
