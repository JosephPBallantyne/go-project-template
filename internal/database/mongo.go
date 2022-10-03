package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (d *Database) InsertMany(collection string, options *options.InsertManyOptions, insert interface{}) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.InsertMany(ctx, insert.([]interface{}), options)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *Database) FindOne(collection string, selector bson.D, options *options.FindOneOptions, output interface{}) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := col.FindOne(ctx, selector, options).Decode(output)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No records found")
	} else if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *Database) FindMany(collection string, selector bson.D, options *options.FindOptions, output interface{}) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := col.Find(ctx, selector, options)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No records found")
	} else if err != nil {
		return err
	}
	err = cur.All(ctx, output)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *Database) UpdateOne(collection string, selector bson.D, options *options.UpdateOptions, update interface{}) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.UpdateOne(ctx, selector, update, options)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *Database) UpdateMany(collection string, selector bson.D, options *options.UpdateOptions, update interface{}) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.UpdateMany(ctx, selector, update, options)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *Database) DeleteOne(collection string, options *options.DeleteOptions, selector bson.D) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.DeleteOne(ctx, selector, options)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *Database) DeleteMany(collection string, options *options.DeleteOptions, selector bson.D) error {
	col := d.client.Database(d.dbName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.DeleteMany(ctx, selector, options)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
