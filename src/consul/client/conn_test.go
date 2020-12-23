package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"keep-learning-go/src/consul/config"
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
