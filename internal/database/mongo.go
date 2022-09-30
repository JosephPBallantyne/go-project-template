package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client *mongo.Client
	dbName string
}

func NewMongoClient(url string, dbName string) (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("client failed")
		return nil, err
	}
	return &Database{
		client,
		dbName,
	}, nil
}

func (d *Database) ConnectClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := d.client.Connect(ctx)
	if err != nil {
		fmt.Println("connection failed")
	}
	// collection := m.Database("test").Collection("user")
	// _, errI := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "one"}, {Key: "value", Value: 1}})
	// fmt.Println(errI)

	return err
}

func (d *Database) PingServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := d.client.Ping(ctx, readpref.Primary())
	return err
}

func (d *Database) InsertOne(collection string, options *options.InsertOneOptions, insert interface{}) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, insert, options)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
