package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

// Config defines the application env vars
type Config struct {
	ConnStr        string `env:"CONN_STR,required"`
	DatabaseName   string `env:"DATABASE_NAME,required"`
	TimeExecImport string `env:"TIME_EXEC_IMPORT,required"`

	EmailUser   string `env:"EMAIL_USER,required"`
	EmailPass   string `env:"EMAIL_PASS,required"`
	EmailServer string `env:"EMAIL_SEVER,required"`
	EmailPort   int    `env:"EMAIL_PORT,required"`
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
