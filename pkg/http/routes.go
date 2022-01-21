package http

import (
	"fmt"
	"net/http"
)

func (h *handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *handler) hello(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("sd")

	response := fmt.Sprint(h.settings.FooSettings.Name, h.settings.FooSettings.Description)
	w.WriteHeader(200)
	w.Write([]byte(response))
}
