package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, error := net.ResolveTCPAddr("tcp4", ":7777")
	if error != nil {
		fmt.Println("Error: %s", error.Error())
		return
	}

	listener, error := net.ListenTCP("tcp4", addr)
	if error != nil {
		fmt.Println("Error: %s", error.Error())
		return
	}
	defer listener.Close()

	for {
		pTCPConn, error := listener.AcceptTCP()
		if error != nil {
			fmt.Println("Error: %s", error.Error())
			continue
		}
		go handler(pTCPConn)
	}
}

func handler(con *net.TCPConn) {
	defer con.Close()
	now := time.Now()
	con.Write([]byte(now.String() + "\n"))
}
