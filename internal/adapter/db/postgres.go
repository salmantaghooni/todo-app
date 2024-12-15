package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"todo-app/internal/domain"
	"todo-app/pkg/config"
)

// NewPostgresDB initializes a new PostgreSQL connection
func NewPostgresDB(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// RunMigrations applies database migrations
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&domain.TodoItem{})
}
