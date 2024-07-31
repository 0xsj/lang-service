package lib

import (
	"net/http"

	"google.golang.org/grpc"
)

type ServiceRegistrar interface {
	RegisterGRPCServices(server *grpc.Server)
	RegisterHTTPHandlers(mux *http.ServeMux)
	Init()
}
