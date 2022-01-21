package main

import (
	"net/http"
	"uberfxsample/pkg/config"
	httphandler "uberfxsample/pkg/http"
	"uberfxsample/pkg/logger"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(logger.ProvideLogger),
		fx.Provide(config.ProvideConfig),
		fx.Provide(http.NewServeMux),
		fx.Invoke(httphandler.New),
		fx.Invoke(httphandler.RegisterHooks),
	).Run()
}

//geri dönen değer diğerleri tarafından kullanılmayacağı zaman invoke
//kullanılacağı zaman provide çağırıyoruz
