package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func DbConnection() *sql.DB {
	//Connection to PostgreSQL
	username := os.Getenv("DB_USER")
	// pasword := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", username, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// --- REMOVED: defer db.Close() ---

	dberr := db.Ping()
	if dberr != nil {
		// Use log.Fatalf here to print the error and exit
		log.Fatalf("Error connecting to database: %v", dberr)
	}

	fmt.Println("DB connected")
	return db
}
