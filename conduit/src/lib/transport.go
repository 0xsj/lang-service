package lib

import (
	"log"
	"net"

	"github.com/0xsj/kakao/conduit/src/config"
	"google.golang.org/grpc"
)

// NewGRPCServer initializes a new gRPC server instance
func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

// StartGRPCServer starts the gRPC server with the provided config
func StartGRPCServer(grpcServer *grpc.Server, cfg *config.Config) {
	// Convert the port to string
	address := net.JoinHostPort("", cfg.AuthService.Port)

	// Create a network listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on address %s: %v", address, err)
	}

	log.Printf("Starting gRPC server on %s\n", address)


	// Start serving gRPC requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
