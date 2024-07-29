package app

import (
	"github.com/0xsj/kakao/auth/src/modules/user"
)

type AppModule struct {
	UserModule *user.UserModule
}


func NewAppModule() *AppModule {
	userModule := user.NewUserModule()

	return &AppModule{
		UserModule: userModule,
	}
}

