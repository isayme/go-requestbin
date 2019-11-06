package conf

import (
	config "github.com/isayme/go-config"
)

var globalConfig Config

// Config service config
type Config struct {
	HTTP struct {
		Port int `json:"port" yaml:"port"`
	} `json:"http" yaml:"http"`

	Logger struct {
		Level string `json:"level" yaml:"level"`
	} `json:"logger" yaml:"logger"`

	Mongo Mongo `json:"mongo" yaml:"mongo"`
}

// Mongo mongo config
type Mongo struct {
	URI string `json:"uri" yaml:"uri"`
}

// Get get config
func Get() *Config {
	config.Parse(&globalConfig)

	return &globalConfig
}
