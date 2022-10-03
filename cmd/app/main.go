package main

import (
	"fmt"
	"os"

	"github.com/josephpballantyne/go-project-template/internal/config"
	"github.com/josephpballantyne/go-project-template/internal/database"
	"github.com/josephpballantyne/go-project-template/internal/http"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	constants, _ := config.InitViper()
	db, _ := database.NewMongoClient(constants.Mongo.URL, constants.Mongo.DBName)
	err := db.ConnectClient()
	if err != nil {
		fmt.Println("db connection failed")
		return err
	}

	us := &database.UserService{Database: db}
	var h http.Handler
	h.UserService = us

	http.StartServer(&h)
	return nil
}
