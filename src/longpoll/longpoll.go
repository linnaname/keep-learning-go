package main

import (
	"fmt"
	"net/http"
	"time"
)

func longOperation(ch chan<- string) {
	time.Sleep(time.Second * 60)
	ch <- "Config changed"
}

func handler(w http.ResponseWriter, _ *http.Request) {
	ch := make(chan string)
	go longOperation(ch)

	select {
	case result := <-ch:
		fmt.Fprint(w, result)
	case <-time.After(time.Second * 20):
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprint(w, "Config not modified.")
	}
	close(ch)
}

func main() {
	http.HandleFunc("/longpolling", handler)
	http.ListenAndServe("localhost:8080", nil)
}
