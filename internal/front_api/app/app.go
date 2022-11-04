package app

import command "github.com/cosmonaut-cat/boardgames_backend/internal/front_api/app/command/game"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddGame command.AddGameHandler
}

type Queries struct{}
