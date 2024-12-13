package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      http.FileServer(http.Dir("./static")),
		ReadTimeout:  10 * time.Second, // Maximum duration for reading the request
		WriteTimeout: 10 * time.Second, // Maximum duration for writing the response
		IdleTimeout:  15 * time.Second, // Maximum duration for idle connections
	}

	log.Println("Server started at :8080")

	// Start the server and handle errors
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
