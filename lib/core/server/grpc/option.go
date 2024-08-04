package grpc

import "go.uber.org/fx"

type Option struct {
	fx.Out
	ServerAddr string 
	ClientAddr string
	// DefaultAddr string
}

func NewOption() Option {
	return Option{
		ServerAddr: ":50051",
		ClientAddr: ":50052",
	}
}

