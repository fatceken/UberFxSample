package app

import (
	"embed"
	"go.uber.org/fx"
	"net/http"
	"uberfxsample/pkg/appsettings"
	"uberfxsample/pkg/configuration"
	httpServer "uberfxsample/pkg/http"
	"uberfxsample/pkg/logger"
	"uberfxsample/pkg/rediscache"
)

//go:embed config.yaml config.*.yaml
var configurationFiles embed.FS

func Create(preInitOpts ...fx.Option) *fx.App {
	return fx.New(CreateCoreOptions(preInitOpts...)...)
}

//geri dönen değer diğerleri tarafından kullanılmayacağı zaman invoke kullanılacağı zaman provide çağırıyoruz

func CreateCoreOptions(preInitOpts ...fx.Option) []fx.Option {
	return []fx.Option{
		fx.Options(preInitOpts...),
		fx.Invoke(
			configureConfigurationOptions,
		),
		configuration.BindConfigToOptions("AppSettings", appsettings.GetType()),
		logger.Module(),
		configuration.Module(),
		appsettings.Module(),
		fx.Provide(http.NewServeMux),
		httpServer.Module(false),
		rediscache.Module(),
	}
}

func configureConfigurationOptions(co *configuration.Options) {
	co.ConfigurationFiles = configurationFiles
}
