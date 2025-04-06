package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Get absolute path to static directory
	staticPath := filepath.Join(".", "static")
	absPath, _ := filepath.Abs(staticPath)
	fmt.Println("Serving files from:", absPath)

	// Verify files exist
	files, err := os.ReadDir(staticPath)
	if err != nil {
		log.Fatalf("Failed to read static directory: %v", err)
	}
	fmt.Println("Available files in static/:")
	for _, file := range files {
		fmt.Println("-", file.Name())
	}

	// Create a file server handler
	fs := http.FileServer(http.Dir("./static"))

	// Configure the HTTP server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      fs, // Use our file server as the handler
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Start the server
	log.Println("Serving files from ./static on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
