package model

import (
	"time"

    "gorm.io/gorm"
)

// User represents a user's entity in the system.
// This struct is used to map the User entity to the database.
type UserModel struct {
    ID           int                `gorm:"primaryKey;autoIncrement"`
    Email        string             `gorm:"type:varchar(255);uniqueIndex;not null"`
    PasswordHash string             `gorm:"type:text;not null"`
    CreatedAt    time.Time          `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time          `gorm:"not null;default:CURRENT_TIMESTAMP"`
    DeletedAt     gorm.DeletedAt    `gorm:"index"`
}

// TableName overrides table name
func (UserModel) TableName() string {
	return "users"
}