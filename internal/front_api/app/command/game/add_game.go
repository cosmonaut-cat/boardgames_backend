package command

import (
	"context"
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/internal/front_api/domain/game"
	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/front_api"
)

type AddGameHandler struct {
	gameRepository game.Repository
}

func NewAddGameHandler(gameRepository game.Repository) AddGameHandler {
	if gameRepository == nil {
		log.Fatalf("Game repository is empty \n")
	}
	return AddGameHandler{gameRepository: gameRepository}
}

func (n AddGameHandler) Handle(ctx context.Context, game *front_api.Game) error {
	err := n.gameRepository.AddGame(ctx, game)

	if err != nil {
		return err
	}

	return nil
}
