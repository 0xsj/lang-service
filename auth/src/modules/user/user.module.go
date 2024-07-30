package user

import "log"

type UserModule struct {
	UserController *UserController
}

func (u *UserModule) Init() {
	log.Println("UserModule initialized successfully")
}

func NewUserModule() *UserModule {
	userController := NewUserController()
	return &UserModule{
		UserController: userController,
	}
}