package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//我怎么觉得这个ResponseWriter和Request的命名像个傻子一样
	handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
