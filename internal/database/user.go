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
	err := s.InsertOne("user", &opt, u)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (s *UserService) GetUser(id int) (map[string]interface{}, error) {
	opt := options.FindOneOptions{}
	query := bson.D{bson.E{Key: "id", Value: id}}
	output := map[string]interface{}{}
	err := s.FindOne("user", query, &opt, &output)
	if err != nil {
		fmt.Println(err)
	}

	return output, err
}
