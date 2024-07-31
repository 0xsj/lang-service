package user

import (
	"log"
	"net/http"
)

type UserModule struct {
	UserController *UserController
	UserService *UserService
}

func (um *UserModule) Init() {
	log.Println("UserModule initialized successfully")
}

func NewUserModule() *UserModule {
	userController := NewUserController()
	return &UserModule{
		UserController: userController,
	}
}

func (um *UserModule) InitRoutes(mux *http.ServeMux) {
	um.UserController.RegisterRoutes(mux)
}