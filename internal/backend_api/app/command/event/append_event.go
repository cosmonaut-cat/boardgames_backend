package command

import (
	"context"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/domain/event"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api"
)

type AppendEventHandler struct {
	eventRepository event.Repository
}

func NewAppendEventHandler(eventRepository event.Repository) AppendEventHandler {
	if eventRepository == nil {
		log.Fatalf("Event repository is empty \n")
	}

	return AppendEventHandler{eventRepository: eventRepository}
}

func (a AppendEventHandler) Handle(ctx context.Context, id string, events []*api.Event) error {
	err := a.eventRepository.AppendEvents(ctx, id, events)

	if err != nil {
		return err
	}

	return nil
}
