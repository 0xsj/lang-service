package server

import "go.uber.org/fx"

type Transport struct {
	fx.Out
}

func NewTransport() Transport {
	return Transport{}
}
