package grpc

import "go.uber.org/fx"

type GRPC struct {
	fx.Out
	ServerAddr string
	ClientAddr string
}

func NewGRPCServer(){}

func Start(){}

func ServerModule() fx.Option{
	return fx.Options()
}