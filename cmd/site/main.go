package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

const timeoutDefault = 10

func main() {
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      http.FileServer(http.Dir("./static")),
		ReadTimeout:  timeoutDefault * time.Second, // Maximum duration for reading the request
		WriteTimeout: timeoutDefault * time.Second, // Maximum duration for writing the response
		IdleTimeout:  timeoutDefault * time.Second, // Maximum duration for idle connections
	}

	log.Println("Server started at :8080")

	// Start the server and handle errors
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start server: %v", err)
	}
}
