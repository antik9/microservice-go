package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config is a configuration from yaml
type Config struct {
	Database struct {
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var (
	Conf Config
)

func init() {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}