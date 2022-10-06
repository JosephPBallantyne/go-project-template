package database

import (
	"context"
	"time"

	"github.com/josephpballantyne/go-project-template/internal/app"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client *mongo.Client
	name   string
}

func ConnectClient(url string, name string) (*Database, error) {
	const op = "NewMongoClient"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Mongo DB failed to connect")
		return nil, &app.Error{Op: op, Err: err}
	}
	log.Info("Mongo DB Client connected")
	return &Database{
		client,
		name,
	}, nil
}

func (d *Database) PingServer() error {
	const op = "Database.PingServer"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := d.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) InsertOne(collection string, options *options.InsertOneOptions, insert interface{}) error {
	const op = "Database.InsertOne"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, insert, options)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) InsertMany(collection string, options *options.InsertManyOptions, insert interface{}) error {
	const op = "Database.InsertMany"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.InsertMany(ctx, insert.([]interface{}), options)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) FindOne(collection string, selector bson.D, options *options.FindOneOptions, output interface{}) error {
	const op = "Database.FindOne"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := col.FindOne(ctx, selector, options).Decode(output)
	if err == mongo.ErrNoDocuments {
		return &app.Error{Op: op, Code: app.ENOTFOUND}
	} else if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) FindMany(collection string, selector bson.D, options *options.FindOptions, output interface{}) error {
	const op = "Database.FindMany"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := col.Find(ctx, selector, options)
	if err == mongo.ErrNoDocuments {
		return &app.Error{Op: op, Code: app.ENOTFOUND}
	} else if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	err = cur.All(ctx, output)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) UpdateOne(collection string, selector bson.D, options *options.UpdateOptions, update interface{}) error {
	const op = "Database.UpdateOne"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.UpdateOne(ctx, selector, update, options)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) UpdateMany(collection string, selector bson.D, options *options.UpdateOptions, update interface{}) error {
	const op = "Database.UpdateMany"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.UpdateMany(ctx, selector, update, options)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) DeleteOne(collection string, options *options.DeleteOptions, selector bson.D) error {
	const op = "Database.DeleteOne"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.DeleteOne(ctx, selector, options)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}

func (d *Database) DeleteMany(collection string, options *options.DeleteOptions, selector bson.D) error {
	const op = "Database.DeleteMany"
	col := d.client.Database(d.name).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.DeleteMany(ctx, selector, options)
	if err != nil {
		return &app.Error{Op: op, Err: err}
	}
	return nil
}
