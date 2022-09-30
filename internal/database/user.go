package database

import (
	"fmt"

	"github.com/josephpballantyne/go-project-template/internal/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	*Database
}

func (s *UserService) CreateUser(u *app.User) error {
	opt := options.InsertOneOptions{}
	insert := bson.D{bson.E{Key: "name", Value: "jeff"}}
	err := s.InsertOne("user", &opt, insert)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
