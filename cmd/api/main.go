package main

import (
	"fmt"
	"log"
	"ordernationn/cmd/db"
	"ordernationn/internal/handlers"
	"ordernationn/internal/service"
	"ordernationn/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := db.DbConnection()
	defer dbConn.Close()
	productStore := store.NewSQLProductRepository(dbConn)
	ProductService := service.NewProductService(productStore)
	productHandler := handlers.NewProductHandler(ProductService)
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.255"})

	router.GET("/api/products/get", productHandler.GetAllProducts, GetClientIP)
	router.GET("/api/product/get", productHandler.GetProductById, GetClientIP)
	// router.StaticFS("/images", gin.Dir("./assets/images/product", true))
	router.GET("/images/*filepath", handlers.ImageHandler)

	log.Println("Server running on :8080")
	log.Fatal(router.Run(":8080"))
}

func GetClientIP(ctx *gin.Context) {
	fmt.Printf("Client IP: %s", ctx.ClientIP())
}
