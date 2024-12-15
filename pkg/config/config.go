package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the entire configuration for the application
type Config struct {
	Server       ServerConfig
	Database     DBConfig
	Storage      StorageConfig
	MessageQueue MQConfig
}

// ServerConfig holds server-related configurations
type ServerConfig struct {
	Port string
}

// DBConfig holds database-related configurations
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// StorageConfig holds AWS S3-related configurations
type StorageConfig struct {
	Endpoint        string
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
}

// MQConfig holds AWS SQS-related configurations
type MQConfig struct {
	Endpoint string
	Region   string
	QueueURL string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DBConfig{
			Host:     getEnv("DB_HOST", "postgres"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "todo_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Storage: StorageConfig{
			Endpoint:        getEnv("S3_ENDPOINT", "http://localhost:4566"),
			Region:          getEnv("S3_REGION", "us-east-1"),
			AccessKeyID:     getEnv("S3_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("S3_SECRET_ACCESS_KEY", "test"),
			Bucket:          getEnv("S3_BUCKET", "uploads"),
		},
		MessageQueue: MQConfig{
			Endpoint: getEnv("SQS_ENDPOINT", "http://localhost:4566"),
			Region:   getEnv("SQS_REGION", "us-east-1"),
			QueueURL: getEnv("SQS_QUEUE_URL", "http://localhost:4566/000000000000/todo-queue"),
		},
	}

	return cfg, nil
}

// getEnv retrieves environment variables or returns a default value
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
