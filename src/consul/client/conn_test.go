package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"keep-learning-go/src/consul/config"
	"net"
	"testing"
)

const CONFIG_PATH = "/Users/goranka/linnana/go/keep-learning-go/config"

func TestNewConn(t *testing.T) {
	config, err := config.NewConfig(CONFIG_PATH, "client")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	bOpt := &BalancerOption{}
	config.Sub("conn").Sub("balancer").Unmarshal(bOpt)
	retryOpt := &RetryOption{}
	config.Sub("conn").Sub("retry").Unmarshal(retryOpt)

	conn, err := NewConn(bOpt, retryOpt)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestGetTarget(t *testing.T) {
	config, err := config.NewConfig(CONFIG_PATH, "client")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	bOpt := &BalancerOption{}
	config.Sub("conn").Sub("balancer").Unmarshal(bOpt)
	target := getTarget(bOpt)
	assert.NotEmpty(t, target)
	println(target)

	println(fmt.Sprintf("%s://%s/%s", "consul", "127.0.0.1:8500", "grpc-go-test"))

}

func TestHostPort(t *testing.T) {
	_, _, err := net.SplitHostPort("127.0.0.1:8500/grpc-go-test")
	if ae, ok := err.(*net.AddrError); ok {
		println(ae.Err)
	} else {
		println("eeer")
	}
}

func TestCheckTarget(t *testing.T) {
	config, _ := config.NewConfig(CONFIG_PATH, "client")
	bOpt := &BalancerOption{}
	config.Sub("conn").Sub("balancer").Unmarshal(bOpt)
	target := getTarget(bOpt)
	_, _, err := net.SplitHostPort(target)
	if ae, ok := err.(*net.AddrError); ok {
		println(ae.Err)
	} else {
		println("eeer")
	}
}
