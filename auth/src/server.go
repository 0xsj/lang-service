package main

import (
	"log"
	"net/http"

	"github.com/0xsj/kakao/auth/src/app"
)

func main() {
	appModule := app.NewAppModule()
	http.HandleFunc("/ping", appModule.UserModule.UserController.HandlePing)

	log.Println("Starting server on :3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}