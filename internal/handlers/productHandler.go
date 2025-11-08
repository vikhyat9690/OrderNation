package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"ordernationn/internal/domain"
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

func (h *ProductHandler) Create(c *gin.Context) {
	var req struct {
		Sku              string  `json:"sku"`
		Name             string  `json:"name"`
		ShortDescription string  `json:"short_description"`
		LongDescription  string  `json:"long_description"`
		Price            float32 `json:"price"`
		SpecialPrice     float64 `json:"special_price"`
		BaseImageUrl     string  `json:"base_image_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json input"})
		return
	}

	product := domain.Product{
		Sku:  req.Sku,
		Name: req.Name,
		Short_Description: sql.NullString{
			String: req.ShortDescription,
			Valid:  req.ShortDescription != "",
		},
		Long_Description: sql.NullString{
			String: req.LongDescription,
			Valid:  req.LongDescription != "",
		},
		Price: req.Price,
		Special_Price: sql.NullFloat64{
			Float64: req.SpecialPrice,
			Valid:   req.SpecialPrice != 0,
		},
		Base_Image_Url: sql.NullString{
			String: req.BaseImageUrl,
			Valid:  req.BaseImageUrl != "",
		},
	}

	createdProduct, err := h.service.Create(product)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdProduct)
}
