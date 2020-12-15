package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

const PATH = "../../img/http.png"

func main() {

	fp, _ := filepath.Abs(PATH)
	fmt.Println(fp)
	fmt.Println(filepath.IsAbs(fp))
	fmt.Println(filepath.Base(PATH))
	fmt.Println(filepath.Split(PATH))
	fmt.Println(filepath.Ext(PATH))

	fmt.Println("=========================")

	fmt.Println("2: ", filepath.Clean("/.../..../////abc/abc"))
	fmt.Println("3: ", filepath.Clean("./1.txt"))

	fmt.Println("=========================")

	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
	fmt.Println(filepath.Dir("../todo.txt"))
	fmt.Println(filepath.Dir(".."))
	fmt.Println(filepath.Dir("."))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))

	fmt.Println("=========================")
	fmt.Println(filepath.Join("a/b", "c"))

	fmt.Println("=========================")
	fmt.Println(filepath.Match(`???`, `abc`)) // true

	fmt.Println("=========================")
	s, err := filepath.Rel(`/a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err)

	fmt.Println("=========================")
	fmt.Println(filepath.SplitList("/a/b/c:/usr/bin"))
	fmt.Println(filepath.SplitList(""))
	fmt.Println(strings.Split("", ":"))

	fmt.Println("=========================")

}
