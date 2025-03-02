package main

import (
	"log"
	"net/http"

	"blog-be/internal/blog/db"
	"blog-be/routes"
)

func main() {
	dsn := "host=localhost user=apple dbname=blog port=5432 sslmode=disable"

	// DB 연결 (필수)
	if err := db.ConnectBoth(dsn); err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}

	r := routes.NewRouter()                    // Initializes the router from the routes package
	log.Fatal(http.ListenAndServe(":8080", r)) // Starts the server on port 8080 and handles incoming HTTP requests
}
