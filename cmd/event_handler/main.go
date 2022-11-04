package main

import (
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/app"
	command "github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/app/command/event"
	"github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/db"
	"github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/services"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/event_handler"
	"github.com/jmoiron/sqlx"
)

var (
	dbClient *sqlx.DB
)

func main() {
	startListener()

	application := newApplication()

	eventService := services.NewEventServiceServer(application)

	event_handler.RegisterEventServicesServer(grpcServer, eventService)

	serveGrpcServer()

	defer dbClient.Close()
}

func newApplication() *app.Application {
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
