package product

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ServiceConfig struct {
	Db            string `yaml:"db"`
	Port          string `yaml:"port"`
	DefaultLimit  int    `yaml:"default_limit"`
	DefaultOffset int    `yaml:"default_offset"`
}

func NewProductServiceConfig(path string) (ServiceConfig, error) {
	var config ServiceConfig

	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
