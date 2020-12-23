package config

import (
	"errors"
	"github.com/spf13/viper"
)

func NewConfig(conf string, configName string) (*viper.Viper, error) {
	if len(configName) == 0 || len(conf) == 0 {
		return nil, errors.New("illegal argument")
	}

	config := viper.New()
	config.SetConfigType("json")
	config.SetConfigName(configName)
	config.AddConfigPath(conf)
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return config, err
}
