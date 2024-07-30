package main

import (
	"github.com/0xsj/kakao/conduit/src/app"
	"github.com/0xsj/kakao/conduit/src/lib"
)

func main() {
	const staticPort = 3333 // Define the static port

	appModule := app.NewAppModule()
	factory := lib.NewFactory(staticPort) // Pass the static port to the factory
	factory.CreateApp(appModule)
	factory.StartMicroservice()
}
