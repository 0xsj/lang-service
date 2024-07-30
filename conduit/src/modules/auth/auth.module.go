package auth

import "log"


type AuthModule struct {
	AuthService *AuthService
	AuthResolver *AuthResolver
}

func (u *AuthModule) Init() {
	log.Println("AuthModule initialized successfully")
}


func NewAuthModule() *AuthModule{
	authService := NewAuthService()

	return &AuthModule{
		AuthService: authService,
	}
}