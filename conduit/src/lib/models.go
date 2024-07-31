package lib

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
)


type TransportType string

const (
	HTTP  TransportType = "http"
	GRPC  TransportType = "grpc"
	RMQ   TransportType = "rmq"
	Kafka TransportType = "kafka"
)

type TransportConfig struct {
	Type        TransportType      // Transport type (HTTP, GRPC, RMQ, Kafka)
	Address     string             // Address to bind the transport
	Port        int                // Port to bind the transport
	ExtraConfig map[string]any     // Extra configuration parameters
}

func (tc TransportConfig) String() string {
	return fmt.Sprintf("Type: %s, Address: %s, Port: %d, ExtraConfig: %+v", tc.Type, tc.Address, tc.Port, tc.ExtraConfig)
}

type ServiceRegistrar interface {
	RegisterGRPCServices(server *grpc.Server)
	RegisterHTTPHandlers(mux *http.ServeMux)
	Init()
}