package http

import (
	"go.uber.org/fx"
	"net/http"

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
	mux    *http.ServeMux
	logger *zap.SugaredLogger
}

func createHandler(s *http.ServeMux, l *zap.SugaredLogger) *handler {
	h := handler{s, l}
	h.registerRoutes()

	return &h
}
