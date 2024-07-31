package app

import (
	"log"

	"github.com/0xsj/kakao/auth/src/modules/user"
)

type AppModule struct {
	UserModule *user.UserModule
}

func (a *AppModule) Init() {
	a.UserModule.Init()
	log.Println("AppModule initialized successfully")
}


func NewAppModule() *AppModule {
	userModule := user.NewUserModule()
	return &AppModule{
		UserModule: userModule,
	}
}
