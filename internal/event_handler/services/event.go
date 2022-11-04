package services

import (
	"context"

	"github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/app"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/event_handler"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventServiceServer struct {
	event_handler.UnimplementedEventServicesServer

	app app.Application
}

func NewEventServiceServer(application *app.Application) *EventServiceServer {
	return &EventServiceServer{app: *application}
}

func (e *EventServiceServer) Append(ctx context.Context, req *event_handler.Event_AppendRequest) (*emptypb.Empty, error) {
	err := e.app.Commands.AppendEvent.Handle(ctx, req.Event)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// func (e *EventServiceServer) Scan(scanRequest *api.Event_ScanRequest, stream api.EventServices_ScanServer) error {
// 	return nil
// }

// func (e *EventServiceServer) Latest(ctx context.Context, scanRequest *api.Event_LatestRequest) (*api.Event, error) {
// 	return nil, nil
// }
