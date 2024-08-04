package interceptor

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module("interceptor", fx.Provide(New, NewValidator, NewPanicRecovery))
}