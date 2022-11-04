package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/event_handler"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/front_api"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MongoGameRepository struct {
	db                 *mongo.Client
	eventHandlerClient event_handler.EventServicesClient
}

func NewMongoGameRepository(db *mongo.Client, eventHandlerClient event_handler.EventServicesClient) *MongoGameRepository {
	if db == nil {
		log.Fatalf("Missing database in game repository\n")
	}

	return &MongoGameRepository{db: db, eventHandlerClient: eventHandlerClient}
}

func (m MongoGameRepository) AddGame(ctx context.Context, game *front_api.Game) error {
	encodedEvent, err := anypb.New(game)

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to marshal game because %s\n", err))
	}

	event := &event_handler.Event{
		EventId:        game.Id,
		EventType:      "game-added",
		EventEntity:    "game",
		EventVersion:   game.Version,
		EventTimestamp: timestamppb.Now(),
		EventPayload:   encodedEvent,
	}

	appendRequest := &event_handler.Event_AppendRequest{
		Event: event,
	}

	_, err = m.eventHandlerClient.Append(ctx, appendRequest)

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to add game because %s\n", err))
	}

	return nil
}
