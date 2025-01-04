package pkg_config

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

func LoadConfig() *Config {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
