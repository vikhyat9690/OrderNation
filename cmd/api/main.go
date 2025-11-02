package main

import (
	"log"
	"net/http"
	"ordernationn/cmd/db"
	"ordernationn/handlers"
	"ordernationn/internal/service"
	"ordernationn/internal/store"
)

func main() {
	dbConn := db.DbConnection()
	defer dbConn.Close()
	s := &store.Store{DB: dbConn}
	svc := &service.OrderService{Store: s}
	handler := handlers.NewOrderHandler(svc)

	http.HandleFunc("/place-order", handler.PlaceOrder())
	
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
