package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/josephpballantyne/go-project-template/api/v1/router"
	"github.com/josephpballantyne/go-project-template/internal/config"
)

func StartServer() {
	r := router.Initialize()
	constants, _ := config.InitViper()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
	fmt.Printf("Starting server on PORT:%s\n", constants.PORT)
	log.Fatal(http.ListenAndServe(":"+constants.PORT, r))
}
