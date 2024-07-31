package app


type AppController struct {
	AppService *AppService
}


func NewAppController(service *AppService) *AppController{
	return &AppController{
		AppService: service,
	}
}
