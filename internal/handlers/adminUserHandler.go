package handlers

import (
	"errors"
	"log"
	"net/http"
	"ordernationn/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminUserHandler struct {
	service service.AdminUserService
}

func NewAdminUserHandler(s service.AdminUserService) *AdminUserHandler {
	return &AdminUserHandler{
		service: s,
	}
}

func (h *AdminUserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json request"})
		return
	}
	var (
		ErrAdminNotFound   = errors.New("admin not found")
		ErrInvalidPassword = errors.New("incorrect password")
	)
	user, err := h.service.AdminLogin(ctx, req.Email, req.Password)
	if err != nil {
		switch err {
		case ErrAdminNotFound:
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "admin not foudn"})
		case ErrInvalidPassword:
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		}
		log.Println(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}
