package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7777")
	if err != nil {
		fmt.Println("resolve ip eror:%v", err)
		return
	}
	con, err := net.DialTCP("tcp", nil /*pLocalTCPAddr*/, tcpAddr)
	if err != nil {
		fmt.Println("dial tcp error %v", err)
		return
	}
	defer con.Close()

	n, err := con.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Println("Write tcp error %v", err)
		return
	}
	fmt.Println("writed: %d\n", n)
	buf, err := ioutil.ReadAll(con)
	r := len(buf)
	fmt.Println(string(buf[:r]))
	fmt.Println("readed: %d\n", r)
}
