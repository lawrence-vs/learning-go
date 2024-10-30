package config

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfig loads configuration from JSON or YAML file.
func LoadConfig(filename string) (*Configuration, error) {
	var config Configuration

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if json.Unmarshal(data, &config) == nil {
		return &config, nil
	}

	if yaml.Unmarshal(data, &config) == nil {
		return &config, nil
	}

	return nil, err
}
