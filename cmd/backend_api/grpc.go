package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcServerListener net.Listener
	grpcServer         *grpc.Server
)

func startListener() {
	opts := []grpc.ServerOption{}

	listener, err := net.Listen("tcp", "0.0.0.0:3030")

	if err != nil {
		log.Fatalf("Failed to listen address %s because %s\n", "0.0.0.0:3030", err)
	}

	grpcServerListener = listener

	grpcServer = grpc.NewServer(opts...)
}

func serveGrpcServer() {
	reflection.Register(grpcServer)

	log.Printf("Server listening at: %s\n", grpcServerListener.Addr())

	if err := grpcServer.Serve(grpcServerListener); err != nil {
		log.Fatalf("Failed to serve becase %s\n", err)
	}
}
