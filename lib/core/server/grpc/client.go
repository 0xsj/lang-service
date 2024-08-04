package grpc

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewGRPCClient(){}

func Close() {}

func ClientModule() fx.Option {
	return fx.Options()
}