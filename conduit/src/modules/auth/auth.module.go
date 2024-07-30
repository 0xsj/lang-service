package auth


type AuthModule struct {
	AuthService *AuthService
	AuthResolver *AuthResolver
}

func NewAuthModule() *AuthModule{
	authService := NewAuthService()

	return &AuthModule{
		AuthService: authService,
	}
}