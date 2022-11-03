package services

import (
	"context"

	"github.com/cosmonaut-cat/boardgames_backend/internal/backend_api/app"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventServiceServer struct {
	api.UnimplementedEventServicesServer

	app app.Application
}

func NewEventServiceServer(application *app.Application) *EventServiceServer {
	return &EventServiceServer{app: *application}
}

func (e *EventServiceServer) Append(ctx context.Context, req *api.Event_AppendRequest) (*emptypb.Empty, error) {
	err := e.app.Commands.AppendEvent.Handle(ctx, req.Id, req.Events)

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
