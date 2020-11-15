package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	Timeout  int `yaml:"timeout"`
	Services []struct {
		Name   string `yaml:"name"`
		URL    string `yaml:"url"`
		Method string `yaml:"method"`
	}
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
