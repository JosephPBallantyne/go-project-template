package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "health check ok")
}
