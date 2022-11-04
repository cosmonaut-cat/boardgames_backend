package main

import (
	"log"

	"github.com/cosmonaut-cat/boardgames_backend/pkg/api/event_handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcClientConn     *grpc.ClientConn
	eventHandlerClient event_handler.EventServicesClient
)

func dialEventGrpcServer() {
	conn, err := grpc.Dial("backend_event_handler:3030", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect event handler server because %s\n", err)
	}

	grpcClientConn = conn

	client := event_handler.NewEventServicesClient(grpcClientConn)

	eventHandlerClient = client
}
