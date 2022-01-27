package http

import (
	"fmt"
	"net/http"
)

func (h *handler) registerRoutes() {
	h.mux.HandleFunc("/hello", h.hello)
	h.mux.HandleFunc("/addToCache", h.addToCache)
	h.mux.HandleFunc("/getFromCache", h.getFromCache)
}

func (h *handler) hello(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("hello world")
	response := fmt.Sprint(h.settings.FooSettings.Name, h.settings.FooSettings.Description)
	w.WriteHeader(200)
	w.Write([]byte(response))
}

//http://localhost:8080/addToCache?key=foo&value=bar
func (h *handler) addToCache(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "The key is missing in query parameter", http.StatusBadRequest)
		return
	}

	value := r.URL.Query().Get("value")
	if value == "" {
		http.Error(w, "The value is missing in query parameter", http.StatusBadRequest)
		return
	}

	err := h.cacheHelper.Add(key, value, 0)
	if err != nil {
		http.Error(w, "Unable to insert", http.StatusInternalServerError)
		h.logger.Error(err.Error())
		return
	}

	response := fmt.Sprintf("Key %v with value %v is inserted", key, value)
	w.WriteHeader(200)
	w.Write([]byte(response))
}

//http://localhost:8080/getFromCache?key=foo
func (h *handler) getFromCache(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "The key is missing in query parameter", http.StatusBadRequest)
		return
	}

	value, err := h.cacheHelper.Get(key)
	if err != nil {
		http.Error(w, "Unable to get", http.StatusInternalServerError)
		h.logger.Error(err.Error())
		return
	}

	response := fmt.Sprintf("The value with key %v is %v ", key, value)
	w.WriteHeader(200)
	w.Write([]byte(response))
}
