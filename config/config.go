package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

// Config defines the application env vars
type Config struct {
	ConnStr string `env:"CONN_STR,required"`
}

// NewConfig will parse the necessary env vars to
// struct Config
func NewConfig() *Config {
	c := new(Config)

	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	return c
}
