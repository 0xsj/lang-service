package lib

import (
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

// Define a common interface for different transports
type Transport interface {
	Start(address string) error
}

// Implement HTTP transport
type HTTPTransport struct {
	Handler http.Handler
}

func (t *HTTPTransport) Start(address string) error {
	fmt.Printf("Starting HTTP server on %s\n", address)
	return http.ListenAndServe(address, t.Handler)
}

// Implement GRPC transport
type GRPCTransport struct {
	Server *grpc.Server
}

func (t *GRPCTransport) Start(address string) error {
	fmt.Printf("Starting gRPC server on %s\n", address)
	// Create a listener and start serving gRPC requests
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	return t.Server.Serve(listener)
}

// StartServer initializes and starts the server based on transport config
func StartServer(config *TransportConfig) error {
	var transport Transport

	switch config.Type {
	case HTTP:
		transport = &HTTPTransport{
			Handler: http.DefaultServeMux, // Replace with actual handler
		}
	case GRPC:
		grpcServer := grpc.NewServer()
		transport = &GRPCTransport{Server: grpcServer}
	case RMQ:
		// Implement RMQTransport similarly
		return fmt.Errorf("RMQ transport not yet implemented")
	case Kafka:
		// Implement KafkaTransport similarly
		return fmt.Errorf("Kafka transport not yet implemented")
	default:
		return fmt.Errorf("unsupported transport type: %s", config.Type)
	}

	if err := transport.Start(config.Address); err != nil {
		return fmt.Errorf("error starting transport: %w", err)
	}

	return nil
}
