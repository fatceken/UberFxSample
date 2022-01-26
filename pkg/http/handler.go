package http

import (
	"context"
	"go.uber.org/fx"
	"net/http"
	"uberfxsample/pkg/appsettings"

	"go.uber.org/zap"
)

var isUnitTesting = false

func Module(isTesting bool) fx.Option {
	isUnitTesting = isTesting
	return fx.Invoke(
		createHandler,
		registerHooks,
	)
}

// handler for http requests
type handler struct {
	mux      *http.ServeMux
	logger   *zap.SugaredLogger
	settings *appsettings.AppSettings
}

func createHandler(s *http.ServeMux, l *zap.SugaredLogger, as *appsettings.AppSettings) *handler {
	h := handler{s, l, as}
	h.registerRoutes()

	return &h
}

func registerHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, s *appsettings.AppSettings, mux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go http.ListenAndServe(s.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				if isUnitTesting {
					return nil
				}
				return logger.Sync()
			},
		},
	)
}
