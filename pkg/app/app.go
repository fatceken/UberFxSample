package app

import (
	"go.uber.org/fx"
	"net/http"
	"uberfxsample/pkg/config"
	httpServer "uberfxsample/pkg/http"
	"uberfxsample/pkg/logger"
)

func Create(preInitOpts ...fx.Option) *fx.App {
	return fx.New(CreateCoreOptions(preInitOpts...)...)
}

//geri dönen değer diğerleri tarafından kullanılmayacağı zaman invoke kullanılacağı zaman provide çağırıyoruz

func CreateCoreOptions(preInitOpts ...fx.Option) []fx.Option {
	return []fx.Option{
		logger.Module(),
		config.Module(),
		fx.Provide(http.NewServeMux),
		httpServer.Module(),
	}
}
