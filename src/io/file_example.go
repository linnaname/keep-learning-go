package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	buf := "Hello, World"
	file, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(buf)); err != nil {
		panic(err)
	}

	fmt.Println(file.Name())
	fmt.Println(path.Base(file.Name()))

	oldpath := file.Name()
	dir := path.Dir(oldpath)
	newpath := filepath.Join(dir, "group-dataId"+".tmp")
	fmt.Println(newpath)
	err = os.Rename(oldpath, newpath)
	fmt.Println(err)
	fmt.Println(file.Name())
	time.Sleep(time.Minute * 2)
}
