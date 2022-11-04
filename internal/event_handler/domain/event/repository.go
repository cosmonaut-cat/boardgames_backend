package event

import (
	"context"

	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/event_handler"
)

const (
	LatestVersion string = "latest_version"
	EventAdded    string = "event_added"
	EventUpdated  string = "event_updated"
)

type Repository interface {
	AppendEvent(ctx context.Context, event *event_handler.Event) error
	// Scan(ctx context.Context, eventId string) error
	Latest(ctx context.Context, eventId string) (*event_handler.Event, error)
}
