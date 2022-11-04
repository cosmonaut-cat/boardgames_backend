package game

import (
	"context"

	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/front_api"
)

type Repository interface {
	AddGame(ctx context.Context, game *front_api.Game) error
}
