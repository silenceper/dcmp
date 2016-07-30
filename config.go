package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

type config struct {
	Listen    string   `yaml:"listen"`
	Endpoints []string `yaml:"endpoints"`
}

func getConfig() *config {
	filename, _ := filepath.Abs("./config/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	var c *config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}
	return c
}
