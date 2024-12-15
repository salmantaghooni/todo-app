package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"todo-app/internal/service"
)

// TodoHandler defines HTTP handlers for Todo operations
type TodoHandler interface {
	CreateTodo(c *gin.Context)
}

// todoHandler is the implementation of TodoHandler
type todoHandler struct {
	service service.TodoService
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(service service.TodoService) TodoHandler {
	return &todoHandler{
		service: service,
	}
}

// CreateTodoRequest represents the request payload for creating a Todo
type CreateTodoRequest struct {
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	FileID      string    `json:"fileId"`
}

// CreateTodo handles the creation of a new TodoItem
func (h *todoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.CreateTodo(req.Description, req.DueDate, req.FileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
