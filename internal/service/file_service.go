package service

import (
	"errors"
	"mime/multipart"
	"strings"

	"todo-app/internal/adapter/storage"
)

// FileService defines methods for file operations
type FileService interface {
	UploadFile(file multipart.File, filename string) (string, error)
}

// fileService is the implementation of FileService
type fileService struct {
	storage storage.S3Client
}

// NewFileService creates a new FileService
func NewFileService(storage storage.S3Client) FileService {
	return &fileService{
		storage: storage,
	}
}

// UploadFile validates and uploads a file to S3
func (s *fileService) UploadFile(file multipart.File, filename string) (string, error) {
	// Validate file type
	if !isAllowedFileType(filename) {
		return "", errors.New("unsupported file type")
	}

	// Upload to S3
	fileID, err := s.storage.UploadFile(file, filename)
	if err != nil {
		return "", err
	}

	return fileID, nil
}

// isAllowedFileType checks if the file has an allowed extension
func isAllowedFileType(filename string) bool {
	allowed := []string{".jpg", ".jpeg", ".png", ".txt"}
	for _, ext := range allowed {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}
	return false
}
