package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server `json:"http_server" yaml:"http_server"`
}

type Server struct {
	Port int `json:"port" yaml:"port"`
}

func NewConfig(filename string) *Config {
	cfg := &Config{}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not open configuration file: %v", err)
	}
	json.Unmarshal(data, cfg)
	return cfg
}
