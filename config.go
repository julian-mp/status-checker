package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ServiceConfig is the config for a single service
type ServiceConfig struct {
	Name    string                 `yaml:"name"`
	URL     string                 `yaml:"url"`
	Method  string                 `yaml:"method"`
	Payload map[string]interface{} `yaml:"payload"`
}

type config struct {
	Timeout  int `yaml:"timeout"`
	Services []ServiceConfig
}

func getConfig(filename string) config {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var cfg config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
