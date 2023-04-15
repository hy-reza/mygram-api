package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name             string    `gorm:"not null" json:"name"`
	Social_media_url string    `gorm:"not null" json:"social_media_url"`
	UserID           string    `gorm:"index;references:ID"`
	CreatedAt        time.Time `gorm:"index"`
	UpdatedAt        time.Time `gorm:"index"`
}

// BeforeCreate hook to set the primary key to a new UUID
func (m *Media) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return
}
