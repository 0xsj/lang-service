package lib

import (
	"fmt"
	"net/http"

	"github.com/0xsj/kakao/conduit/src/config"
)

// StartServer starts an HTTP server with the given handler and configuration.
func StartServer(handler http.Handler, config *config.Config, port int) {
	if port == 0 {
		port = 3000
	}
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("🚀 Server is running on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
