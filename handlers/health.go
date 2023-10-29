package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Health(mux chi.Router) {
	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// handlers default to returning HTTP 200 OK if nothing else is set.
		// 🤪
	})
}
