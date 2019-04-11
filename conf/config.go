package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Path string
	}
}

func LoadConfig(configFilePath string) (*Config, error) {
	var config Config
	source, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
