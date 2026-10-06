package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShutDownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("HTTP", &config); err != nil {
		return nil, fmt.Errorf("Problem creating server cfg :%w", err)
	}

	return &config, nil
}

func ConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		fmt.Printf("get server sfg error : %v", err)
		panic("err")
	}
	return config
}
