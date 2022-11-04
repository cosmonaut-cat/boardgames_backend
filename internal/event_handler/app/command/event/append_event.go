package command

import (
	"context"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/domain/event"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/event_handler"
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

func (a AppendEventHandler) Handle(ctx context.Context, event *event_handler.Event) error {
	err := a.eventRepository.AppendEvent(ctx, event)

	if err != nil {
		return err
	}

	return nil
}
