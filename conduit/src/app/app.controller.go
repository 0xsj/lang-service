package app

import "net/http"


type AppController struct {
	AppService *AppService
}


func NewAppController(service *AppService) *AppController{
	return &AppController{
		AppService: service,
	}
}

func (ac *AppController) GetData(w http.ResponseWriter, r *http.Request){
	data := ac.AppService.GetData()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}