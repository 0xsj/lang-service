package user

import (
	"log"
	"net/http"
)

type UserModule struct {
	UserController *UserController
}

func (um *UserModule) Init() {
	log.Println("UserModule initialized successfully")
	um.RegisterRoutes()
}

func NewUserModule() *UserModule {
	userController := NewUserController()
	return &UserModule{
		UserController: userController,
	}
}

func (um *UserModule) RegisterRoutes(){
	http.HandleFunc("/ping", um.UserController.HandlePing)
}