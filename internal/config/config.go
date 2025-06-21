package config

import (
	"fmt"
	"os"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/.config/notekeeper", home)
}
