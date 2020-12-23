package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const CONFIG_PATH = "/Users/goranka/linnana/go/keep-learning-go/config"

func TestNewConfig(t *testing.T) {
	config, err := NewConfig(CONFIG_PATH, "server")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Get("register.port"))
}

func TestConfigNotExist(t *testing.T) {
	config, err := NewConfig(CONFIG_PATH, "what")
	assert.Error(t, err)
	assert.Nil(t, config)
}
