package database_test

import (
	"testing"

	"github.com/josephpballantyne/go-project-template/internal/app"
	"github.com/josephpballantyne/go-project-template/internal/database"
	"github.com/josephpballantyne/go-project-template/internal/mock"
)

func TestCreateUser(t *testing.T) {
	u := &app.User{
		ID:      1,
		Name:    "tester",
		Address: "home",
	}

	db := &mock.Database{}
	us := database.NewUserService(db)
	err := us.CreateUser(u)

	app.Ok(t, err)
	app.Equals(t, db.InsertOneInvoked, true)
}

func TestFindUser(t *testing.T) {

	db := &mock.Database{}
	us := database.NewUserService(db)
	_, err := us.GetUser(1)

	app.Ok(t, err)
	app.Equals(t, db.FindOneInvoked, true)
}
