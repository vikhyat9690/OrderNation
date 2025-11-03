package handlers

import (
	"log"
	"net/http"
	"ordernationn/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

// Get All Products List
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

// Get product by id
func (h *ProductHandler) GetProductById(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id == 0 {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "incorrect id type passed"})
		return
	}
	product, err := h.service.GetById(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error in fetching product with given id"})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}
