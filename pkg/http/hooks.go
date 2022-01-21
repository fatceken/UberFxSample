package http

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"uberfxsample/pkg/config"
)

func RegisterHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, cfg *config.Config, mux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go http.ListenAndServe(cfg.ApplicationConfig.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
