package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":7778")
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 256)
		n, udpAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: %s", err.Error())
			return
		}
		fmt.Println("readed: %d", n)
		n, err = conn.WriteToUDP(buf, udpAddr)
		if err != nil {
			fmt.Println("Error: %s", err.Error())
			return
		}
		fmt.Println("writed: %d", n)
	}
}
