package main

import (
	"github.com/spf13/viper"
	"os"
)

func NewConfig(conf string) (*viper.Viper, error) {
	config := viper.New()
	fp, err := os.Open(conf)
	if err != nil {
		return nil, err
	}
	if err := config.ReadConfig(fp); err != nil {
		return nil, err
	}
	return config, nil
}
