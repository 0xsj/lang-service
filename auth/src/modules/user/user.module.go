package user

type UserModule struct {
	UserController *UserController
}

func NewUserModule() *UserModule {
	userController := NewUserController()
	return &UserModule{
		UserController: userController,
	}
}