package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Username  string    `gorm:"type:varchar(100);not null;uniqueIndex" `
	Email     string    `gorm:"type:varchar(100);not null;uniqueIndex" `
	Password  string    `gorm:"type:varchar(100);not null" `
	Age       int       `gorm:"type:int;not null" `
	Photos    []Photo   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments  []Comment `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}

// BeforeCreate hook to set the primary key to a new UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
