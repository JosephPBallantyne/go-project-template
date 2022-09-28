package router

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/josephpballantyne/go-project-template/api/v1/handlers"
)

func Initialize() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api", handlers.Routes())
	})
	return r
}
