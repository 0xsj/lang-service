package main

import (
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

	transportConfig := lib.TransportConfig{
		Type:        lib.HTTP,
		Address:     cfg.AuthService.Port,
	}

	if err := lib.StartServer(&transportConfig); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	appModule := app.NewAppModule()
	factory := lib.NewFactory(cfg.AuthService.Port)
	factory.CreateApp(appModule)
	factory.StartMicroservice()
}
