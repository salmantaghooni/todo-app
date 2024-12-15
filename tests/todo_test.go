package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"todo-app/internal/domain"
	"todo-app/internal/service"
)

// MockSQSClient mocks the SQSClient interface
type MockSQSClient struct {
	mock.Mock
}

func (m *MockSQSClient) SendMessage(message interface{}) error {
	args := m.Called(message)
	return args.Error(0)
}

// MockTodoRepository mocks the TodoRepository interface
type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(todo *domain.TodoItem) error {
	args := m.Called(todo)
	return args.Error(0)
}

func TestCreateTodo(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	mockSQS := new(MockSQSClient)

	todoService := service.NewTodoService(mockRepo, mockSQS)

	description := "Test Todo"
	dueDate := time.Now().Add(24 * time.Hour)
	fileID := "file123"

	mockRepo.On("Create", mock.Anything).Return(nil)
	mockSQS.On("SendMessage", mock.Anything).Return(nil)

	todo, err := todoService.CreateTodo(description, dueDate, fileID)

	assert.NoError(t, err)
	assert.Equal(t, description, todo.Description)
	assert.Equal(t, fileID, todo.FileID)

	mockRepo.AssertExpectations(t)
	mockSQS.AssertExpectations(t)
}
