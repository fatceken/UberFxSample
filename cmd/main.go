package main

import (
	"io/ioutil"
	"net/http"
	"uberfxsample/config"
	"uberfxsample/httphandler"

	"go.uber.org/zap"

	"gopkg.in/yaml.v2"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slogger := logger.Sugar()

	mux := http.NewServeMux()
	httphandler.New(mux, slogger)

	conf := &config.Config{}
	data, err := ioutil.ReadFile("config/base.yaml")

	if err != nil {
		slogger.Error(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)

	if err != nil {
		slogger.Error(err)
	}

	http.ListenAndServe(conf.ApplicationConfig.Address, mux)

}
