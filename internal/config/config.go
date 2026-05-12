package config

import (
	"go-slim.dev/env"
)

type DatabaseConfig struct {
	Name     string `env:"POSTGRES_DB"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"POSTGRES_HOST"`
	MaxConns int32  `env:"DB_MAX_CONNS"`
}

type ApplicationConfig struct {
	Port int `env:"APP_PORT"`
}

type Config struct {
	Database    DatabaseConfig
	Application ApplicationConfig
}

func LoadConfig() (*Config, error) {
	if err := env.Init(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Signed("", "").Fill(&cfg); err != nil {
		return nil, err
	}
	env.Lock()
	return &cfg, nil
}
