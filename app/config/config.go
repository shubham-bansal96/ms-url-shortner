package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Log         Logging `yaml:"Logging"`
	MSName      string  `yaml:"MSName"`
	Environment string  `yaml:"Environment"`
}

type Logging struct {
	LogLevel string `yaml:"LogLevel"`
}

var Config *Configuration

func Initialize() {
	ymlConfig, err := ioutil.ReadFile(FileName)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err.Error())
	}
	if err := yaml.Unmarshal(ymlConfig, &Config); err != nil {
		log.Fatalf("Error Unmarshelling config file: %v", err.Error())
	}
}

// helper method to set mock config
func SetConfig(config *Configuration) {
	Config = config
}
