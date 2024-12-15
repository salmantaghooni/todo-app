package service

import (
	"time"

	"todo-app/internal/adapter/mq"
	"todo-app/internal/domain"
	"todo-app/internal/repository"
)

// TodoService defines methods for Todo operations
type TodoService interface {
	CreateTodo(description string, dueDate time.Time, fileID string) (*domain.TodoItem, error)
}

// todoService is the implementation of TodoService
type todoService struct {
	repo repository.TodoRepository
	mq   mq.SQSClient
}

// NewTodoService creates a new TodoService
func NewTodoService(repo repository.TodoRepository, mq mq.SQSClient) TodoService {
	return &todoService{
		repo: repo,
		mq:   mq,
	}
}

// CreateTodo creates a new TodoItem and sends it to SQS
func (s *todoService) CreateTodo(description string, dueDate time.Time, fileID string) (*domain.TodoItem, error) {
	todo := &domain.TodoItem{
		Description: description,
		DueDate:     dueDate,
		FileID:      fileID,
	}

	// Save to DB
	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}

	// Send to SQS
	if err := s.mq.SendMessage(todo); err != nil {
		return nil, err
	}

	return todo, nil
}
