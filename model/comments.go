package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    string    `gorm:"index;references:ID"`
	PhotoID   string    `gorm:"index;references:ID"`
	Message   string    `gorm:"not null" json:"message" `
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}

// BeforeCreate hook to set the primary key to a new UUID
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
