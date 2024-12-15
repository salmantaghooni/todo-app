package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo-app/internal/service"
)

// FileHandler defines HTTP handlers for file operations
type FileHandler interface {
	Upload(c *gin.Context)
}

// fileHandler is the implementation of FileHandler
type fileHandler struct {
	service service.FileService
}

// NewFileHandler creates a new FileHandler
func NewFileHandler(service service.FileService) FileHandler {
	return &fileHandler{
		service: service,
	}
}

// Upload handles file upload requests
func (h *fileHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	fileID, err := h.service.UploadFile(file, header.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fileId": fileID})
}
