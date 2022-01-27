package http

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"net/http"
	"testing"
	"uberfxsample/pkg/appsettings"
	"uberfxsample/pkg/rediscache"
)

func TestModule(t *testing.T) {
	app := fxtest.New(t,
		fx.Provide(http.NewServeMux),
		fx.Provide(func() *zap.SugaredLogger {
			logger, _ := zap.NewProduction()
			return logger.Sugar()
		}),
		fx.Provide(func() *appsettings.AppSettings {
			return &appsettings.AppSettings{
				Address: "8080",
			}
		}),
		fx.Provide(func() *rediscache.Helper {
			return &rediscache.Helper{}
		}),
		Module(true),
	)

	app.RequireStart()
	app.RequireStop()
}
