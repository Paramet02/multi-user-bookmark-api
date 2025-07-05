package domain

import (
	"time"
)

// User represents a user in the system.
// not have no tags or gorm annotations
type User struct {
	// Entity fields
	ID           int
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time // pointer to time.Time to allow null value
}