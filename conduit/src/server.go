package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0xsj/kakao/conduit/src/app"
)

func main() {
	appModule := app.NewAppModule()
	fmt.Println(appModule)
	log.Println("starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}