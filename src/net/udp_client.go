package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7778")
	if err != nil {
		fmt.Println("Error ResolveUDPAddr: %s", err.Error())
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		fmt.Println("Error DialUDP: %s", err.Error())
		return
	}

	n, err := conn.Write([]byte("你好啊！！！"))
	if err != nil {
		fmt.Println("Error WriteToUDP: %s", err.Error())
		return
	}

	fmt.Println("writed: %d", n)
	buf := make([]byte, 1024)
	n, _, err = conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error ReadFromUDP: %s", err.Error())
		return
	}

	fmt.Println("readed: %d  %s", n, string(buf[:n]))
}
