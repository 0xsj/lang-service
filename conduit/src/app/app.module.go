package app

import (
	"log"

	"github.com/0xsj/kakao/conduit/src/modules/auth"
)

type AppModule struct {
	AuthModule *auth.AuthModule
}

func (a *AppModule) Init() {
	a.AuthModule.Init()
	log.Println("AppModule initialized successfully")
}



func NewAppModule() *AppModule {
	authModule := auth.NewAuthModule()
	return &AppModule{
		AuthModule: authModule,
	}
}