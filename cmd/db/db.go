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
		log.Fatal(err)
	}
	defer db.Close()

	dberr := db.Ping()
	if dberr != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	// s := &store.Store{DB: db}
	// svc := &service.OrderService{Store: s}

	fmt.Println("DB connected")
	// _ = svc
	return db
}
