package main

import (
	"github.com/0xsj/kakao/lib/core/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		server.NewModule(),
	)	

	app.Run()
}