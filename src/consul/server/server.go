package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"keep-learning-go/src/consul/api"
	"keep-learning-go/src/consul/config"
	"keep-learning-go/src/consul/service"
	"net"
)

func main() {
	register := NewConsulRegister()
	config, err := config.NewConfig("./config", "server")
	if err != nil {
		panic(err)
	}
	config.Sub("register").Unmarshal(register)
	if err := register.Register(); err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	api.RegisterServiceServer(server, new(service.CalService))
	grpc_health_v1.RegisterHealthServer(server, &service.HealthService{})

	address, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", config.Get("register.port")))
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
