package main

import (
	"blog-be/internal/blog/db"
	"blog-be/routes" // Importing the routes package
	"log"
	"net/http"
)

func main() {
	// Initialize database connection
	db.Connect() // Connects to the PostgreSQL database

	// Initialize and start the HTTP server
	r := routes.NewRouter()                    // Initializes the router from the routes package
	log.Fatal(http.ListenAndServe(":8080", r)) // Starts the server on port 8080 and handles incoming HTTP requests
}
