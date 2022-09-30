package http

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/josephpballantyne/go-project-template/internal/app"
)

type Handler struct {
	UserService app.UserService
}

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &app.User{
			ID:      1,
			Name:    "Aaa",
			Address: "123",
		}
		h.UserService.CreateUser(u)
		render.JSON(w, r, "user created")
	}
}

func (h *Handler) GetHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "health check ok")
	}
}
