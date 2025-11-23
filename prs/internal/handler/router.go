package handler

import (
	"prs/internal/service"

	"github.com/go-chi/chi/v5"
)

func RegisterRouters(r chi.Router, s service.PRService) {
	handler := NewPRSHandler(s)

	r.Route("/team", func (r chi.Router) {
		r.Post("/add", handler.AddTeam)
		r.Get("/get", handler.GetTeam)
	})

	r.Route("/users", func (r chi.Router) {
		r.Post("/setIsActive", handler.UserSetIsActive)
	})
}