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


// Factory struct now handles ports as int or string
type Factory struct {
	config *config.Config
	port   int
}

// NewFactory creates a new Factory instance with the provided port
func NewFactory(port interface{}) *Factory {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var intPort int

	switch v := port.(type) {
	case int:
		intPort = v
	case string:
		var err error
		intPort, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Invalid port value: %v", err)
		}
	default:
		log.Fatalf("Invalid port type: %T", v)
	}

	return &Factory{
		config: cfg,
		port:   intPort,
	}
}

// GetPort returns the port as an int
func (f *Factory) GetPort() int {
	return f.port
}

// CreateApp initializes the application module
func (f *Factory) CreateApp(module AppModule) {
	module.Init()
	log.Println("App initialized with module:", module)
}

// StartMicroservice starts the microservice server
func (f *Factory) StartMicroservice() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Microservice is running"))
	})

	address := net.JoinHostPort("", strconv.Itoa(f.GetPort())) 
	log.Printf("Starting microservice on port %d\n", f.GetPort())
	log.Fatal(http.ListenAndServe(address, nil))
}
