package app

import "github.com/0xsj/kakao/auth/src/modules/user"

type AppController struct {
	UserController *user.UserController
}

func NewAppController(userController *user.UserController) *AppController{
	return &AppController{
		UserController: userController,
	}
}
