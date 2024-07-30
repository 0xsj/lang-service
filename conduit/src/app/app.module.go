package app

import (
	"github.com/0xsj/kakao/conduit/src/modules/auth"
)

type AppModule struct {
	AuthModule *auth.AuthModule
}


func NewAppModule() *AppModule {
	authModule := auth.NewAuthModule()
	return &AppModule{
		AuthModule: authModule,
	}
}