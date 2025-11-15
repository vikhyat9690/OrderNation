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
	adminUserStore := store.NewSQLAdminUserRepository(dbConn)
	ProductService := service.NewProductService(productStore)
	AdminUserService := service.NewAdminUserService(adminUserStore)
	productHandler := handlers.NewProductHandler(ProductService)
	adminUserHandler := handlers.NewAdminUserHandler(AdminUserService)
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.255"})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET(
		"/api/products/get",
		GetClientIP,
		middleware.ApiKeyAuthMiddlware(apiKey),
		productHandler.GetAllProducts,
	)

	router.GET(
		"/api/product/get",
		GetClientIP,
		middleware.ApiKeyAuthMiddlware(apiKey),
		productHandler.GetProductById,
	)

	router.POST(
		"/api/product/add",
		GetClientIP,
		middleware.ApiKeyAuthMiddlware(apiKey),
		productHandler.Create,
	)
	// router.StaticFS("/images", gin.Dir("./assets/images/product", true))
	router.GET("/images/*filepath", handlers.ImageHandler)

	router.POST(
		"/api/admin/login",
		GetClientIP,
		middleware.ApiKeyAuthMiddlware(apiKey),
		adminUserHandler.Login,
	)

	log.Println("Server running on :8080")
	log.Fatal(router.Run(":8080"))
}

func GetClientIP(ctx *gin.Context) {
	fmt.Printf("Client IP: %s", ctx.ClientIP())
}
