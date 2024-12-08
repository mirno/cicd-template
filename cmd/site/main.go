package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Server started at :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
