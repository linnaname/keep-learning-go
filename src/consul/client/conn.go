package main

import (
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"strings"
)

func NewConn(bOpt *BalancerOption, rOpt *RetryOption) (*grpc.ClientConn, error) {
	resolver.Register(NewConsulResolver())
	target := getTarget(bOpt)
	conn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(rOpt.WaitBetween)),
				grpc_retry.WithMax(rOpt.Retries),
				grpc_retry.WithPerRetryTimeout(rOpt.Timeout),
				grpc_retry.WithCodes(rOpt.Codes...),
			),
		),
	)
	return conn, err
}

func getTarget(bOpt *BalancerOption) string {
	builder := strings.Builder{}
	builder.WriteString("consul://")
	builder.WriteString(bOpt.Consul)
	builder.WriteString("/")
	builder.WriteString(bOpt.Service)
	return builder.String()
}
