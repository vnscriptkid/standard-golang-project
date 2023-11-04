package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type pinger interface {
	Ping(ctx context.Context) error
}

func Health(mux chi.Router, p pinger) {
	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// handlers default to returning HTTP 200 OK if nothing else is set.
		if err := p.Ping(r.Context()); err != nil {
			// 502 Bad Gateway: signalling that an underlying resource is at fault
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	})
}
