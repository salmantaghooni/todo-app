package repository

import (
	"todo-app/internal/domain"

	"gorm.io/gorm"
)

// TodoRepository defines methods to interact with TodoItems in the DB
type TodoRepository interface {
	Create(todo *domain.TodoItem) error
}

// todoRepository is the implementation of TodoRepository
type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new TodoRepository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

// Create inserts a new TodoItem into the database
func (r *todoRepository) Create(todo *domain.TodoItem) error {
	if err := r.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
