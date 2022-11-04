package main

import (
	"context"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/internal/front_api/app"
	command "github.com/cosmonaut-cat/boardgames_backend/internal/front_api/app/command/game"
	"github.com/cosmonaut-cat/boardgames_backend/internal/front_api/db"
	"github.com/cosmonaut-cat/boardgames_backend/internal/front_api/services"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/front_api"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbClient *mongo.Client
)

func main() {
	startListener()

	dialEventGrpcServer()

	application := newApplication()

	gameService := services.NewGameServiceServer(application)

	front_api.RegisterGameServicesServer(grpcServer, gameService)

	serveGrpcServer()

	defer grpcClientConn.Close()

	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect database because %s\n", err)
		}
	}()

}

func newApplication() *app.Application {
	dbClient, err := db.NewMongoConnection()

	if err != nil {
		log.Fatalf("Failed to connect database because %s\n", err)
	}

	gameRepository := db.NewMongoGameRepository(dbClient, eventHandlerClient)

	return &app.Application{
		Commands: app.Commands{
			AddGame: command.NewAddGameHandler(gameRepository),
		},
	}
}
