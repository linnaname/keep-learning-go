package main

import (
	"github.com/stretchr/testify/assert"
	"keep-learning-go/src/consul/config"
	"testing"
)

const CONFIG_PATH = "/Users/goranka/linnana/go/keep-learning-go/config"

func TestLocalIP(t *testing.T) {
	ip := localIP()
	assert.NotEmpty(t, ip)
	println(ip)
}

func TestRegister(t *testing.T) {
	config, err := config.NewConfig(CONFIG_PATH, "server")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	register := NewConsulRegister()
	assert.NotNil(t, register)
	config.Sub("register").Unmarshal(register)
	err = register.Register()
	assert.NoError(t, err)
}
