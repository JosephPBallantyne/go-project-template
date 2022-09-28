package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/josephpballantyne/go-project-template/internal/services"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", services.GetHealth)
	return r
}
