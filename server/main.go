package main

import (
	"github.com/mullayam/go-esp8266-iot/internal"
	"log"
	"net/http"
)

func main() {

	mux := routes.RegisterRoutes()
	// Start the HTTP server
	port := ":8080"
	log.Printf("Server starting on %s...", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
