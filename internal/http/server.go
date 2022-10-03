package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/josephpballantyne/go-project-template/internal/config"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router *chi.Mux
}

func StartServer(h *Handler) *server {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)
	r.Route("/v1/api", func(r chi.Router) {
		r.Get("/health", h.GetHealth())
		r.Post("/user", h.CreateUser())
		r.Get("/user/{id}", h.GetUser())
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	constants, _ := config.InitViper()
	log.WithFields(log.Fields{"PORT": constants.PORT}).Info("Server starting")
	log.Fatal(http.ListenAndServe(":"+constants.PORT, r))
	return &server{
		router: r,
	}
}
