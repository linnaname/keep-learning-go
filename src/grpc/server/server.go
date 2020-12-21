package main

import (
	"google.golang.org/grpc"
	"keep-learning-go/src/grpc/message"
	"net"
)

func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
