package database

import (
	"github.com/josephpballantyne/go-project-template/internal/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	*Database
}

func (s *UserService) CreateUser(u *app.User) error {
	const op = "UserService.CreateUser"
	opt := options.InsertOneOptions{}
	err := s.InsertOne("user", &opt, u)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (s *UserService) GetUser(id int) (map[string]interface{}, error) {
	const op = "UserService.GetUser"
	opt := options.FindOneOptions{}
	query := bson.D{bson.E{Key: "id", Value: id}}
	output := map[string]interface{}{}
	err := s.FindOne("user", query, &opt, &output)
	if err != nil {
		return output, &app.Error{Op: op, Err: err}
	}
	return output, nil
}
