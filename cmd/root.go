package cmd

import (
	"github.com/josephpballantyne/go-project-template/cmd/db"
	"github.com/josephpballantyne/go-project-template/cmd/server"
)

func Execute() {
	db.NewConnection()
	server.StartServer()
}
