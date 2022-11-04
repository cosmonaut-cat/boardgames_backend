package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection() (*mongo.Client, error) {
	uri := os.Getenv("READ_DATABASE_URI")

	if uri == "" {
		log.Fatalln("Failed to connect to read_db, you must set your READ_DATABASE_URI env variable")
	}

	clientOptions := []*options.ClientOptions{
		options.Client().ApplyURI(uri),
	}

	client, err := mongo.Connect(context.TODO(), clientOptions...)

	if err != nil {
		return nil, err
	}

	return client, nil
}
