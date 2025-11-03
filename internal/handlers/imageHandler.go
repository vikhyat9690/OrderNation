package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const ImageRoot = "assets/upload/images"

func ImageHandler(c *gin.Context) {
	filePathParam := c.Param("filepath")

	fullPath := filepath.Join(ImageRoot, filePathParam)
	log.Println("Attempting to access:", fullPath)

	safePath := filepath.Clean(fullPath)

	fileInfo, err := os.Stat(safePath)
	log.Println("os.Stat error:", err)
	if os.IsNotExist(err) || (err == nil && fileInfo.IsDir()) {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	if err != nil {
		log.Println("Error accessing file:", err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.File(safePath)
}
