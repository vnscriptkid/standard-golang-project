package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"canvas/views"
)

func FrontPage(mux chi.Router) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// return 500
		http.Error(w, "error", http.StatusInternalServerError)

		_ = views.FrontPage().Render(w)
	})
}
