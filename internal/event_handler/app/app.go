package app

import command "github.com/cosmonaut-cat/boardgames_backend/internal/event_handler/app/command/event"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AppendEvent command.AppendEventHandler
}

type Queries struct {
}
