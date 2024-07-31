package main

import (
	"fmt"
	"log"

	"github.com/0xsj/kakao/auth/src/app"
	"github.com/0xsj/kakao/auth/src/config"
	"github.com/0xsj/kakao/auth/src/lib"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	fmt.Println(cfg.AuthService.Port)
	

	transportConfig := lib.TransportConfig{
		Type:    lib.HTTP,
		Port:    3000,
	}

	appModule := app.NewAppModule()
	app := lib.NewFactory()
	app.CreateApp(appModule)



	if err := lib.StartServer(&transportConfig); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
