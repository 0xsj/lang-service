package app

type AppService struct {}

func NewAppService() *AppService  {
	return &AppService{}
}

func (as *AppService) GetData() string {
	return `{"message": "Hello from gateway service"}`
}