package db

import (
	"context"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/front_api"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoGameRepository struct {
	db *mongo.Client
}

func NewMongoGameRepository(db *mongo.Client) *MongoGameRepository {
	if db == nil {
		log.Fatalf("Missing database in game repository\n")
	}

	return &MongoGameRepository{db: db}
}

func (m MongoGameRepository) AddGame(ctx context.Context, game *front_api.Game) error {
	return nil
}
