package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import for PostgreSQL driver
)

var DB *sql.DB

func ConnectDatabase() {
	// Database connection parameters
	host := "localhost"
	port := 5432
	user := "your_username"
	password := "your_password"
	dbname := "your_database"

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to open database connection: %v", err)
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Database connected successfully!")
}