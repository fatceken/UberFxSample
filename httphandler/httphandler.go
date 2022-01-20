package httphandler

import (
	"net/http"

	"go.uber.org/zap"
)

// Handler for http requests
type Handler struct {
	mux    *http.ServeMux
	logger *zap.SugaredLogger
}

// New http handler
func New(s *http.ServeMux, l *zap.SugaredLogger) *Handler {
	h := Handler{s, l}
	h.registerRoutes()

	return &h
}

// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("hello called")

	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}
