package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect establishes a connection to the database
func Connect() {
	var err error
	DB, err = sql.Open("postgres", "user=your_user dbname=blog sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
}
