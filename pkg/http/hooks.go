package http

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"uberfxsample/pkg/appsettings"
)

func registerHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, s *appsettings.AppSettings, mux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go http.ListenAndServe(s.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
