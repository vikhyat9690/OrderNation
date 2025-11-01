package main

import (
	"database/sql"
	"fmt"
	"log"
	"ordernationn/internal/service"
	"ordernationn/internal/store"
	"os"

	_ "github.com/lib/pq"
)

func main() {
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
	s := &store.Store{DB: db}
	svc := &service.OrderService{Store: s}

	fmt.Println("System initialized. Ready For processes")
	_ = svc
}
