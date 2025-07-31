package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("no key")
	}
	return &Config{Key: key}
}
