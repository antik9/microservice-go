package hello

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config is a configuration from yaml
type Config struct {
	Web struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
		Name string `yaml:"name"`
	} `yaml:"web"`
	Log struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"log"`
}

var (
	ServerConfig Config
)

func init() {
	yamlFile, err := ioutil.ReadFile("configs/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &ServerConfig)
	if err != nil {
		log.Fatal(err)
	}
}
