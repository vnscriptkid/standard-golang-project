package server

import (
	"canvas/handlers"
	"canvas/model"
	"context"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)

	handlers.FrontPage(s.mux)
	handlers.NewsletterSignup(s.mux, &signupperMock{})
	handlers.NewsletterThanks(s.mux)
}

type signupperMock struct{}

func (s signupperMock) SignupForNewsletter(ctx context.Context, email model.Email) (string, error) {
	return "", nil
}
