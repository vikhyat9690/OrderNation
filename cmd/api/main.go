package main

import (
	"fmt"
	"log"
	"ordernationn/cmd/db"
	"ordernationn/internal/handlers"
	"ordernationn/internal/middleware"
	"ordernationn/internal/service"
	"ordernationn/internal/store"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dbConn := db.DbConnection()
	defer dbConn.Close()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not found in environment")
	}
	productStore := store.NewSQLProductRepository(dbConn)
	ProductService := service.NewProductService(productStore)
	productHandler := handlers.NewProductHandler(ProductService)
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.255"})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // your Next.js frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/api/products/get", productHandler.GetAllProducts, GetClientIP)
	router.GET("/api/product/get", productHandler.GetProductById, GetClientIP)
	router.POST("/api/product/add", GetClientIP, middleware.ApiKeyAuthMiddlware(apiKey), productHandler.Create)
	// router.StaticFS("/images", gin.Dir("./assets/images/product", true))
	router.GET("/images/*filepath", handlers.ImageHandler)

	log.Println("Server running on :8080")
	log.Fatal(router.Run(":8080"))
}

func GetClientIP(ctx *gin.Context) {
	fmt.Printf("Client IP: %s", ctx.ClientIP())
}
