package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// If no name is provided, default to "World"
	if name == "" {
		name = "World"
	}

	w.Header().Set("Content-Type", "text/plain")
	// Write the response
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Set up the HTTP server and route
	http.HandleFunc("/hello", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}

}
