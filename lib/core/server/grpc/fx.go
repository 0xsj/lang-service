package grpc

import "go.uber.org/fx"


func NewModule() fx.Option {
	return fx.Module("grpc", fx.Provide(), ServerModule(), ClientModule())
}