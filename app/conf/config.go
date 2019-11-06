package conf

import (
	"encoding/json"
	"io/ioutil"

	logger "github.com/isayme/go-logger"
)

var config Config

// Config service config
type Config struct {
	HTTP struct {
		Port int `json:"port"`
	} `json:"http"`

	Logger struct {
		Level string `json:"level"`
	} `json:"logger"`

	Mongo Mongo `json:"mongo"`
}

// Mongo mongo config
type Mongo struct {
	URI string `json:"uri"`
}

// SetPath set config file path
func SetPath(path string) {
	logger.Debugf("config file path: %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}

// Get get config
func Get() *Config {
	return &config
}
