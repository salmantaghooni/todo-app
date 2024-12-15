package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todo-app/internal/adapter/db"
	"todo-app/internal/adapter/mq"
	"todo-app/internal/adapter/storage"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"
	"todo-app/pkg/config"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize the database
	dbConn, err := db.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Run database migrations
	if err := db.RunMigrations(dbConn); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	// Initialize S3 storage
	s3Client, err := storage.NewS3Client(cfg.Storage)
	if err != nil {
		log.Fatalf("Could not connect to S3: %v", err)
	}

	// Initialize SQS
	sqsClient, err := mq.NewSQSClient(cfg.MessageQueue)
	if err != nil {
		log.Fatalf("Could not connect to SQS: %v", err)
	}

	// Initialize repositories and services
	todoRepo := repository.NewTodoRepository(dbConn)
	todoService := service.NewTodoService(todoRepo, sqsClient)
	fileService := service.NewFileService(s3Client)

	// Initialize handlers
	todoHandler := handler.NewTodoHandler(todoService)
	fileHandler := handler.NewFileHandler(fileService)

	// Setup Gin router
	router := gin.Default()

	// Define routes
	router.POST("/upload", fileHandler.Upload)
	router.POST("/todo", todoHandler.CreateTodo)

	// Start the server
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
