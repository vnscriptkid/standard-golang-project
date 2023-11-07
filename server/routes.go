package server

import (
	"canvas/handlers"
	"canvas/model"
	"context"

	"github.com/go-chi/chi/middleware"
)

func (s *Server) setupRoutes() {
	s.mux.Use(handlers.AddMetrics(s.metrics))
	handlers.Health(s.mux, s.database)

	handlers.FrontPage(s.mux)
	handlers.NewsletterSignup(s.mux, s.database, s.queue)
	handlers.NewsletterThanks(s.mux)
	handlers.NewsletterConfirm(s.mux, s.database, s.queue)
	handlers.NewsletterConfirmed(s.mux)

	metricsAuth := middleware.BasicAuth("metrics", map[string]string{"prometheus": s.metricsPassword})
	handlers.Metrics(s.mux.With(metricsAuth), s.metrics)
}

type signupperMock struct{}

func (s signupperMock) SignupForNewsletter(ctx context.Context, email model.Email) (string, error) {
	return "", nil
}
