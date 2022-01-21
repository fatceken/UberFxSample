package http

import "net/http"

func (h *handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *handler) hello(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("hello called")

	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}
