package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct{
		Port int64
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	f, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	}

	return c
}
