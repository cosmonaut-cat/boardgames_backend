package event

import (
	"context"

	"github.com/cosmonaut-cat/boardgames_backend/pkg/api"
)

const (
	LatestVersion string = "latest_version"
	EventAdded    string = "event_added"
	EventUpdated  string = "event_updated"
)

type Repository interface {
	AppendEvents(ctx context.Context, eventId string, events []*api.Event) error
	// Scan(ctx context.Context, eventId string) error
	Latest(ctx context.Context, eventId string) (*api.Event, error)
}
