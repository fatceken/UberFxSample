package http

import (
	"go.uber.org/fx"
	"net/http"
	"uberfxsample/pkg/appsettings"

	"go.uber.org/zap"
)

func Module() fx.Option {
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
