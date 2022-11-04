package services

import (
	"context"

	"github.com/cosmonaut-cat/boardgames_backend/internal/front_api/app"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/front_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GameServiceServer struct {
	front_api.UnimplementedGameServicesServer

	app app.Application
}

func NewGameServiceServer(application *app.Application) *GameServiceServer {
	return &GameServiceServer{app: *application}
}

func (g *GameServiceServer) AddGame(ctx context.Context, req *front_api.Game_AddOrUpdateGameRequest) (*emptypb.Empty, error) {
	err := g.app.Commands.AddGame.Handle(ctx, req.Game)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
