package main

import (
	"os"
)

const (
	DefaultServerAddress  = ":8080"
	DefaultMetricsAddress = ":9108"
)

type Config struct {
	ServerAddress  string
	MetricsAddress string
}

func NewConfig() *Config {
	var config Config

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		config.ServerAddress = DefaultServerAddress
	} else {
		config.ServerAddress = serverAddress
	}

	metricsAddress := os.Getenv("METRICS_ADDRESS")
	if metricsAddress == "" {
		config.MetricsAddress = DefaultMetricsAddress
	} else {
		config.MetricsAddress = metricsAddress
	}

	return &config
}
