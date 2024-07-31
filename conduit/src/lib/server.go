package lib

import (
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Transport interface {
	Start(address string) error
}

type HTTPTransport struct {
	Handler http.Handler
}

func (t *HTTPTransport) Start(address string) error {
	fmt.Printf("Starting HTTP server on %s\n", address)
	return http.ListenAndServe(address, t.Handler)
}

type GRPCTransport struct {
	Server *grpc.Server
}

func (t *GRPCTransport) Start(address string) error {
	fmt.Printf("Starting gRPC server on %s\n", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	return t.Server.Serve(listener)
}

func StartServer(config *TransportConfig) error {
	var transport Transport

	switch config.Type {
	case HTTP:
		transport = &HTTPTransport{
			Handler: http.DefaultServeMux,
		}
	case GRPC:
		grpcServer := grpc.NewServer()
		reflection.Register(grpcServer)
		transport = &GRPCTransport{Server: grpcServer}
	case RMQ:
		return fmt.Errorf("RMQ transport not yet implemented")
	default:
		return fmt.Errorf("unsupported transport type: %s", config.Type)
	}

	address := fmt.Sprintf(":%d", config.Port) 
	if err := transport.Start(address); err != nil {
		return fmt.Errorf("error starting transport: %w", err)
	}

	return nil
}
