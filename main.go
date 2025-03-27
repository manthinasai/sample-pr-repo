package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler will respond to the root URL with a welcome message.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go Web Server using Gorilla Mux!")
}

// AboutHandler will respond to the /about URL with some information.
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple web server built using Go and Gorilla Mux.")
}

// Start the server and define routes.
func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define routes and their handlers
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/about", AboutHandler).Methods("GET")

	// Start the server
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
