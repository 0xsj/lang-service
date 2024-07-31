package lib

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/0xsj/kakao/auth/src/config"
)

type AppModule interface {
	Init()
}

type Factory struct {
	config *config.Config
	port   int
}

func NewFactory() *Factory {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return &Factory{
		config: cfg,
	}
}

func (f *Factory) GetPort() int {
	return f.port
}

func (f *Factory) CreateApp(module AppModule) {
	module.Init()
	log.Println("App initialized with module:", module)
}

func (f *Factory) StartMicroservice() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Microservice is running"))
	})

	address := net.JoinHostPort("", strconv.Itoa(f.GetPort())) 
	log.Printf("Starting microservice on port %d\n", f.GetPort())
	log.Fatal(http.ListenAndServe(address, nil))
}

// Add a method to start the app on a specific port
func (f *Factory) StartAppServer(port int) {
	address := net.JoinHostPort("", strconv.Itoa(port))
	log.Printf("Starting server on port %d\n", port)
	
	// Assuming this is an HTTP server; adjust as needed for your use case
	http.ListenAndServe(address, nil)
}
