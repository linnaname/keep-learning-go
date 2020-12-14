package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := "192.168.1.97"
	ip := net.ParseIP(name)
	println(ip.String())
	println(ip.DefaultMask().String())
	println(ip.Mask(ip.DefaultMask()).String())

	domain := "www.baidu.com"
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s IP: %s Network: %s Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	for _, n := range ns {
		fmt.Fprintln(os.Stdout, n)
	}

	port, err := net.LookupPort("tcp", "telnet")
	fmt.Println(port)
}
