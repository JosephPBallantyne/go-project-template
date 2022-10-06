package database

import (
	"github.com/josephpballantyne/go-project-template/internal/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Operations interface {
	InsertOne(collection string, options *options.InsertOneOptions, insert interface{}) error
	FindOne(collection string, selector bson.D, options *options.FindOneOptions, output interface{}) error
}

type UserService struct {
	Operations
}

func NewUserService(o Operations) *UserService {
	return &UserService{o}
}

func (us *UserService) CreateUser(u *app.User) error {
	const op = "UserService.CreateUser"
	opt := options.InsertOneOptions{}
	err := us.InsertOne("user", &opt, u)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (us *UserService) GetUser(id int) (map[string]interface{}, error) {
	const op = "UserService.GetUser"
	opt := options.FindOneOptions{}
	query := bson.D{bson.E{Key: "id", Value: id}}
	output := map[string]interface{}{}
	err := us.FindOne("user", query, &opt, &output)
	if err != nil {
		return output, &app.Error{Op: op, Err: err}
	}
	return output, nil
}
