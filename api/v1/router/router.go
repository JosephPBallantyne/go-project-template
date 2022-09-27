package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/josephpballantyne/go-project-template/api/v1/handlers"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", handlers.GetHealth)
	return r
}
