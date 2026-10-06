package core_pgx_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Database string        `envconfig:"DB" required:"true"`
	TimeOut  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("POSTGRES", &config); err != nil {
		return nil, fmt.Errorf("retrieve config form env : %w", err)
	}
	return &config, nil
}

func ConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		fmt.Printf("error postgres cfg :%w", err)
		panic("Error during postgres cfg build from env")
	}
	return config
}
