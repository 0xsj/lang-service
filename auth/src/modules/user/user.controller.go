package user

import (
	"encoding/json"
	"net/http"
)

type UserController struct {}

func NewUserController() *UserController {
	return &UserController{}
}
func (uc *UserController) HandlePing(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "pong"}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}