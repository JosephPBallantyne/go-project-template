package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/josephpballantyne/go-project-template/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Constants config.Constants
	Database  *mongo.Client
}

func NewConnection() {
	constants, err := config.InitViper()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(constants.Mongo.URL).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("test").Collection("user")
	_, errI := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 3.14159}})
	fmt.Println(errI)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
