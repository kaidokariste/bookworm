package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// Represents database server and credentials
type Config struct {
	DBUrl string
	AuthDatabase string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("C:/find/in/config.go", &c); err != nil {
		log.Fatal(err)
	}
}

