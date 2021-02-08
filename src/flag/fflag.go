package main

import (
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
)

func main() {
	nameSrvAdders := flag.String("n", "", "name server address,spe with ; when more than one")
	if !flag.Parsed() {
		flag.Parse()
	}
	fmt.Println(*nameSrvAdders)

	if u, err := user.Current(); err == nil {
		fmt.Println("用户ID: " + u.Uid)
		fmt.Println("主组ID: " + u.Gid)
		fmt.Println("用户名: " + u.Username)
		fmt.Println("主组名: " + u.Name)
		fmt.Println("家目录: " + u.HomeDir)
	}

	fmt.Println(filepath.Base("/Users/goranka/linnana/go/gdiamond/server/etc/etc.toml"))
}
