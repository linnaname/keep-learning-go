package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

var modifyMarkCache sync.Map

func TestName(t *testing.T) {
	s := make([]interface{}, 1, 10)
	println(len(s))
	println(s[0] == nil)

	v, ok := modifyMarkCache.Load("key")
	println(ok)
	println(v == nil)
	act, loader := modifyMarkCache.LoadOrStore("test", "true")
	println(loader)
	println(act == "true")
	println(act == nil)
	//println(getCurrentDirectory())
	modifyMarkCache.Store("test", nil)
	fmt.Println(modifyMarkCache.Load("test"))
	//println(execPath())

	//ex, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//exPath := filepath.Dir(ex)
	//fmt.Println(exPath)

}

func getCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}
