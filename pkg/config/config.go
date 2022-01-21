package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
)

func ProvideConfig(logger *zap.SugaredLogger) *Config {
	conf := Config{}

	data, err := ioutil.ReadFile("pkg/config/base.yaml")
	if err != nil {
		logger.Error(err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		logger.Error(err)
	}

	return &conf
}
