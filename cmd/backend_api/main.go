package main

import (
	"context"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/app"
	command "github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/app/command/event"
	"github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/db"
	"github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/services"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api"
	"github.com/jmoiron/sqlx"
)

var (
	dbClient *sqlx.DB
)

func main() {
	ctx := context.Background()

	startListener()

	application := newApplication(ctx)

	eventService := services.NewEventServiceServer(application)

	api.RegisterEventServicesServer(grpcServer, eventService)

	serveGrpcServer()

	defer dbClient.Close()
}

func newApplication(ctx context.Context) *app.Application {
	dbClient, err := db.NewMariaDBConnection()

	if err != nil {
		log.Fatalf("Failed to connect database because %s\n", err)
	}

	eventsRepository := db.NewMariaDBEventRepository(dbClient)

	return &app.Application{
		Commands: app.Commands{
			AppendEvent: command.NewAppendEventHandler(eventsRepository),
		},
	}

}
