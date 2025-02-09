package server

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	*mongo.Client
	Db *mongo.Database
}

func NewMongoClient() (*MongoClient, error) {
	// TODO: CE-32 get from env
	mongoUri := "mongodb://localhost:27018"

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	// TODO: CE-32 get from env
	db := client.Database("go-sse")
	db.CreateCollection(context.TODO(), "sessions")
	return &MongoClient{
		Client: client,
		Db:     db,
	}, nil
}
