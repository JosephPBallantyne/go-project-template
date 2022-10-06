package mock

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	InsertOneInvoked bool
	FindOneInvoked   bool
}

func (d *Database) InsertOne(collection string, options *options.InsertOneOptions, insert interface{}) error {
	d.InsertOneInvoked = true
	return nil
}

func (d *Database) FindOne(collection string, selector bson.D, options *options.FindOneOptions, output interface{}) error {
	d.FindOneInvoked = true
	return nil
}
