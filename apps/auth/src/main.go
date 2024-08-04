package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "auth service")
	}

	http.HandleFunc("/", handler)

	port := ":3002"
	fmt.Printf("server is lisening on port%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}