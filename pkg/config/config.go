package config

import (
	"go.uber.org/fx"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
)

func Module() fx.Option {
	return fx.Provide(
		createConfig,
	)
}

func createConfig(logger *zap.SugaredLogger) *Config {
	conf := Config{}

	data, err := ioutil.ReadFile("pkg/app/base.yaml")
	if err != nil {
		logger.Error(err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		logger.Error(err)
	}

	return &conf
}
