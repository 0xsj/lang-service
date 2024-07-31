package app

type AppService struct {}

func NewAppService() *AppService {
	return &AppService{}
}

func (as *AppService) GetPing() string {
	return `${"message": "Hello from auth service"}`
}