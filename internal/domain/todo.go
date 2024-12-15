package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TodoItem represents a task with optional file attachment
type TodoItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	FileID      string    `json:"fileId,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// BeforeCreate hook to set UUID before inserting into DB
func (t *TodoItem) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
