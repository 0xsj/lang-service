package server

import (
	"github.com/0xsj/kakao/lib/core/server/grpc"
	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("server",
		fx.Options( grpc.NewModule()))
}
