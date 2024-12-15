package tests

import (
	"bytes"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"todo-app/internal/adapter/storage"
	"todo-app/internal/service"
)

// MockS3Client mocks the S3Client interface
type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) UploadFile(file multipart.File, filename string) (string, error) {
	args := m.Called(file, filename)
	return args.String(0), args.Error(1)
}

func TestUploadFile(t *testing.T) {
	mockS3 := new(MockS3Client)
	fileService := service.NewFileService(mockS3)

	// Create a dummy file
	fileContent := []byte("This is a test file")
	buf := bytes.NewBuffer(fileContent)
	file := storage.NonClosableFile{Reader: buf}
	filename := "test.txt"

	mockS3.On("UploadFile", mock.Anything, filename).Return("file123", nil)

	fileID, err := fileService.UploadFile(file, filename)

	assert.NoError(t, err)
	assert.Equal(t, "file123", fileID)

	mockS3.AssertExpectations(t)
}

// NonClosableFile is a helper type to mock multipart.File
type NonClosableFile struct {
	*bytes.Buffer
}

func (f NonClosableFile) Close() error {
	return nil
}
