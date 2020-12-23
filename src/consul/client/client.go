package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"keep-learning-go/src/consul/api"
	"keep-learning-go/src/consul/config"
	"math/rand"
	"time"
)

type BalancerOption struct {
	Consul  string
	Service string
}

type RetryOption struct {
	WaitBetween time.Duration
	Retries     uint
	Timeout     time.Duration
	Codes       []codes.Code
}

func main() {
	config, err := config.NewConfig("./config", "client")
	if err != nil {
		panic(err)
	}

	bOpt := &BalancerOption{}
	config.Sub("conn").Sub("balancer").Unmarshal(bOpt)

	retryOpt := &RetryOption{}
	config.Sub("conn").Sub("retry").Unmarshal(retryOpt)

	conn, err := NewConn(bOpt, retryOpt)
	if err != nil {
		fmt.Printf("dial failed. err: [%v]\n", err)
		return
	}
	defer conn.Close()

	client := api.NewServiceClient(conn)
	time.Sleep(time.Millisecond * 200)

	for i := 0; i < 10; i++ {
		request := &api.Request{
			A: int64(rand.Intn(1000)),
			B: int64(rand.Intn(1000)),
		}

		response, err := client.Add(context.Background(), request)
		if err != nil {
			fmt.Println("client.Add err:", err)
			continue
		}
		fmt.Println("client.Add response:", response)

	}

}
